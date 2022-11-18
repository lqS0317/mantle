/* Imports: External */
import { spawn, spawnSync } from 'child_process'

import { DeployFunction } from 'hardhat-deploy/dist/types'
import { hexStringEquals, awaitCondition } from '@mantleio/core-utils'
import { ethers } from 'ethers'
import { HttpNetworkConfig } from 'hardhat/types'

import { names } from '../src/address-names'
import {
  deployAndVerifyAndThen,
  getContractFromArtifact,
} from '../src/deploy-utils'

// run one terminal with `npx hardhat node --no-deploy`
// in another run `npx hardhat deploy --tags DataAvailability`

/* disable autmoning with
curl -H "Content-Type: application/json" -X POST --data \
        '{"id":1337,"jsonrpc":"2.0","method":"evm_setAutomine","params":[false]}' \
        http://localhost:8545

      and then mine with


      curl -H "Content-Type: application/json" -X POST --data \
        '{"id":1337,"jsonrpc":"2.0","method":"evm_setIntervalMining","params":[false]}' \
        http://localhost:8545

        run these in a third terminal when the node is up and running

*/
const deployFn: DeployFunction = async (hre) => {
  const { deployer } = await hre.getNamedAccounts()

  const owner = hre.deployConfig.bvmAddressManagerOwner
  const l1BitAddress = hre.deployConfig.l1BitAddress
  // deploy impl

  console.log('call EigenLayr DataLayr deployer')

  console.log('deploying and verifying EigenDataLayr')
  const rpc = 'http://127.0.0.1:8545' // HACK get from hre
  console.log({ rpc })
  const deployerPrivateKey =
    '0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80' // HACK get from hre
  console.log({ deployerPrivateKey })
  console.log('deploying dlsm')
  spawnSync(
    `echo "yo" && pwd && cd ../../datalayr-mantle/contracts/eignlayr-contracts && mkdir -p data && forge script ./script/Deployer.s.sol --hardhat --rpc-url ${rpc} --private-key ${deployerPrivateKey} --broadcast -vvvv`,
    [],
    { shell: true, stdio: 'inherit' }
  )
  console.log('deployed dlsm')

  // now get addresses from out directory
  // const _dlsm = require('../../datalayr-mantle/contracts/eignlayr-contracts/data/dlsm.json')
  const _dlsm = '0x2E2Ed0Cfd3AD2f1d34481277b3204d807Ca2F8c2' // HACK get from file above
  console.log({ _dlsm })

  // get the EigenDataLayrRollup contract bytecode
  const contractBytecode =
    // eslint-disable-next-line @typescript-eslint/no-var-requires
    require('../contracts/data-availability/out/EigenDataLayrRollup.sol/EigenDataLayrRollup.json')
      .bytecode.object
  const contractABI =
    // eslint-disable-next-line @typescript-eslint/no-var-requires
    require('../contracts/data-availability/out/EigenDataLayrRollup.sol/EigenDataLayrRollup.json').abi
  // console.log(contractBytecode)

  const factory = new ethers.ContractFactory(contractABI, contractBytecode)
  console.log('deploying EigenDataLayrRollup ...')

  // the datalayr rollup contract has the following params:
  // address _sequencer,
  // IERC20 _stakeToken,
  // uint256 _neededStake,
  // IDataLayrServiceManager _dlsm

  const _sequencer = hre.deployConfig.bvmSequencerAddress
  const _stakeToken = hre.deployConfig.l1BitAddress
  const _neededStake = hre.deployConfig.requiredDAStakeAmount
  const datalayrRollup = await factory.deploy(
    _sequencer,
    _stakeToken,
    _neededStake,
    _dlsm
  )

  console.log('deploy EigenDataLayrRollup success at ', datalayrRollup.address)

  // deploy proxy
  const Impl_EigenDataLayrRollup = await getContractFromArtifact(
    hre,
    names.managed.contracts.EigenDataLayrRollup,
    {
      iface: 'EigenDataLayrRollup',
      signerOrProvider: deployer,
    }
  )
}

// This is kept during an upgrade. So no upgrade tag.
deployFn.tags = ['DataAvailability', 'upgrade']

export default deployFn
