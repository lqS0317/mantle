import { task} from "hardhat/config"

/**
 * hh queryReward --tssRewardContract 0xf5e5b77dd4040f5f4c2c1ac8ab18968ef79fd6fe --network rinkeby
 */
task("queryReward", "获取Consult")
  .addParam("tssRewardContract", "tssRewardContract合约地址")
  .setAction(async (taskArgs, hre) => {
    const tssRewardContract = await hre.ethers.getContractFactory('TssRewardContract')
    const slidingWindowOracleContracts = await tssRewardContract.attach(taskArgs.slidingwindoworacle)
    const queryReward= await slidingWindowOracleContracts.queryReward()
    console.log("queryReward->", queryReward)
  });


module.exports = {}
