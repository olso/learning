const ProgressbarTableToken = artifacts.require("ProgressbarTableToken");
const BigNumber = web3.BigNumber;

require("chai")
  .use(require("chai-bignumber")(BigNumber))
  .should();

contract("ProgressbarTableToken", ([_, creator]) => {
  beforeEach(async () => {
    this.pbtt = await ProgressbarTableToken.new({ from: creator });
  });

  it("has a name", async () => {
    const name = await this.pbtt.name();
    assert.equal(name, "ProgressbarTableToken");
  });

  it("has a symbol", async () => {
    const symbol = await this.pbtt.symbol();
    assert.equal(symbol, "PBTT");
  });

  it("has 18 decimals", async () => {
    const decimals = await this.pbtt.decimals();
    decimals.should.be.bignumber.equal(18);
  });

  it("assigns the initial total supply to the creator", async () => {
    const totalSupply = await this.pbtt.totalSupply();
    const creatorBalance = await this.pbtt.balanceOf(creator);

    creatorBalance.should.be.bignumber.equal(totalSupply);

    // const receipt = await web3.eth.getTransactionReceipt(this.pbtt.transactionHash);
  });
});
