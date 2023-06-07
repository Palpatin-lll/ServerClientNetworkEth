// node call_contract_set.js HelloSol *******
var username = process.argv[2];
var contract_to_deploy = process.argv[3];
var unlock_password = process.argv[4];

console.log('Contract script: ' + contract_to_deploy);
var abi_path = 'build/' + contract_to_deploy + '.abi';
var abi;
var myCoinbase;

var fs = require("fs");
var Web3 = require('web3');
const net = require('net');
const web3 = new Web3(new Web3.providers.IpcProvider(`/home/${username}/node1/geth.ipc`, net));

var version = web3.version;
console.log('Web3 version: ' + version);

var contract_address = fs.readFileSync(contract_to_deploy + '.address').toString();
console.log('contract_address: ' + contract_address);

web3.eth.getCoinbase()
.then(coinbase => {
  myCoinbase = coinbase;
  console.log('coinbase: ' + coinbase);
  return coinbase;
})
.then(function (account) {
  return web3.eth.personal.unlockAccount(account, unlock_password, 600)
})
.then(function (unlocked) {
  console.log('Unlocked: ' + unlocked);
  abi = JSON.parse(fs.readFileSync(abi_path), 'utf8');
  var myContract = new web3.eth.Contract(abi, contract_address);

  myContract.methods.withdrawAll().send({from: myCoinbase })
.once('setValue transactionHash', (hash) => {
  console.log('hash: ' + hash);
})
.on('setValue confirmation', (confNumber) => {
  console.log('confNumber: ' + confNumber);
})
.on('receipt', (receipt) => {
  console.log(JSON.stringify(receipt, undefined, 2));
})
})
.catch(function (error) {
  console.error(error);
});
