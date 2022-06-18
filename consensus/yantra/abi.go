package yantra

const validatorsAbi = `[
	{
		"inputs": [
			{
				"internalType": "address[]",
				"name": "_validators",
				"type": "address[]"
			},
			{
				"internalType": "uint256[]",
				"name": "_caps",
				"type": "uint256[]"
			},
			{
				"internalType": "address",
				"name": "_firstOwner",
				"type": "address"
			},
			{
				"internalType": "uint256",
				"name": "_minValidatorCap",
				"type": "uint256"
			},
			{
				"internalType": "uint256",
				"name": "_maxValidatorNumber",
				"type": "uint256"
			},
			{
				"internalType": "uint256",
				"name": "_validatorWithdrawDelay",
				"type": "uint256"
			},
			{
				"internalType": "uint256",
				"name": "_delegatorWithdrawDelay",
				"type": "uint256"
			}
		],
		"stateMutability": "nonpayable",
		"type": "constructor"
	},
	{
		"anonymous": false,
		"inputs": [
			{
				"indexed": false,
				"internalType": "address",
				"name": "_voter",
				"type": "address"
			},
			{
				"indexed": false,
				"internalType": "address",
				"name": "_candidate",
				"type": "address"
			},
			{
				"indexed": false,
				"internalType": "uint256",
				"name": "_cap",
				"type": "uint256"
			}
		],
		"name": "Delegate",
		"type": "event"
	},
	{
		"anonymous": false,
		"inputs": [
			{
				"indexed": true,
				"internalType": "address",
				"name": "previousOwner",
				"type": "address"
			},
			{
				"indexed": true,
				"internalType": "address",
				"name": "newOwner",
				"type": "address"
			}
		],
		"name": "OwnershipTransferred",
		"type": "event"
	},
	{
		"anonymous": false,
		"inputs": [
			{
				"indexed": false,
				"internalType": "address",
				"name": "_owner",
				"type": "address"
			},
			{
				"indexed": false,
				"internalType": "address",
				"name": "_candidate",
				"type": "address"
			},
			{
				"indexed": false,
				"internalType": "uint256",
				"name": "_cap",
				"type": "uint256"
			}
		],
		"name": "Propose",
		"type": "event"
	},
	{
		"anonymous": false,
		"inputs": [
			{
				"indexed": false,
				"internalType": "address",
				"name": "_owner",
				"type": "address"
			},
			{
				"indexed": false,
				"internalType": "address",
				"name": "_candidate",
				"type": "address"
			}
		],
		"name": "Resign",
		"type": "event"
	},
	{
		"anonymous": false,
		"inputs": [
			{
				"indexed": false,
				"internalType": "address",
				"name": "_voter",
				"type": "address"
			},
			{
				"indexed": false,
				"internalType": "address",
				"name": "_candidate",
				"type": "address"
			},
			{
				"indexed": false,
				"internalType": "uint256",
				"name": "_cap",
				"type": "uint256"
			}
		],
		"name": "Undelegate",
		"type": "event"
	},
	{
		"anonymous": false,
		"inputs": [
			{
				"indexed": false,
				"internalType": "address",
				"name": "_owner",
				"type": "address"
			},
			{
				"indexed": false,
				"internalType": "uint256",
				"name": "_blockNumber",
				"type": "uint256"
			},
			{
				"indexed": false,
				"internalType": "uint256",
				"name": "_cap",
				"type": "uint256"
			}
		],
		"name": "Withdraw",
		"type": "event"
	},
	{
		"inputs": [
			{
				"internalType": "address",
				"name": "_candidate",
				"type": "address"
			}
		],
		"name": "delegate",
		"outputs": [],
		"stateMutability": "payable",
		"type": "function"
	},
	{
		"inputs": [],
		"name": "delegatorWithdrawDelay",
		"outputs": [
			{
				"internalType": "uint256",
				"name": "",
				"type": "uint256"
			}
		],
		"stateMutability": "view",
		"type": "function"
	},
	{
		"inputs": [
			{
				"internalType": "address",
				"name": "_candidate",
				"type": "address"
			},
			{
				"internalType": "address",
				"name": "_delegator",
				"type": "address"
			}
		],
		"name": "getDelegatorCap",
		"outputs": [
			{
				"internalType": "uint256",
				"name": "",
				"type": "uint256"
			}
		],
		"stateMutability": "view",
		"type": "function"
	},
	{
		"inputs": [
			{
				"internalType": "address",
				"name": "_candidate",
				"type": "address"
			}
		],
		"name": "getDelegators",
		"outputs": [
			{
				"internalType": "address[]",
				"name": "",
				"type": "address[]"
			}
		],
		"stateMutability": "view",
		"type": "function"
	},
	{
		"inputs": [],
		"name": "getRewards",
		"outputs": [
			{
				"internalType": "uint256[]",
				"name": "",
				"type": "uint256[]"
			}
		],
		"stateMutability": "view",
		"type": "function"
	},
	{
		"inputs": [
			{
				"internalType": "address",
				"name": "_candidate",
				"type": "address"
			}
		],
		"name": "getValidatorCap",
		"outputs": [
			{
				"internalType": "uint256",
				"name": "",
				"type": "uint256"
			}
		],
		"stateMutability": "view",
		"type": "function"
	},
	{
		"inputs": [
			{
				"internalType": "address",
				"name": "_candidate",
				"type": "address"
			}
		],
		"name": "getValidatorOwner",
		"outputs": [
			{
				"internalType": "address",
				"name": "",
				"type": "address"
			}
		],
		"stateMutability": "view",
		"type": "function"
	},
	{
		"inputs": [],
		"name": "getValidators",
		"outputs": [
			{
				"internalType": "address[]",
				"name": "",
				"type": "address[]"
			}
		],
		"stateMutability": "view",
		"type": "function"
	},
	{
		"inputs": [],
		"name": "getWithdrawBlockNumbers",
		"outputs": [
			{
				"internalType": "uint256[]",
				"name": "",
				"type": "uint256[]"
			}
		],
		"stateMutability": "view",
		"type": "function"
	},
	{
		"inputs": [
			{
				"internalType": "uint256",
				"name": "_blockNumber",
				"type": "uint256"
			}
		],
		"name": "getWithdrawCap",
		"outputs": [
			{
				"internalType": "uint256",
				"name": "",
				"type": "uint256"
			}
		],
		"stateMutability": "view",
		"type": "function"
	},
	{
		"inputs": [
			{
				"internalType": "address",
				"name": "_candidate",
				"type": "address"
			}
		],
		"name": "isValidator",
		"outputs": [
			{
				"internalType": "bool",
				"name": "",
				"type": "bool"
			}
		],
		"stateMutability": "view",
		"type": "function"
	},
	{
		"inputs": [],
		"name": "maxValidatorNumber",
		"outputs": [
			{
				"internalType": "uint256",
				"name": "",
				"type": "uint256"
			}
		],
		"stateMutability": "view",
		"type": "function"
	},
	{
		"inputs": [],
		"name": "minValidatorCap",
		"outputs": [
			{
				"internalType": "uint256",
				"name": "",
				"type": "uint256"
			}
		],
		"stateMutability": "view",
		"type": "function"
	},
	{
		"inputs": [],
		"name": "owner",
		"outputs": [
			{
				"internalType": "address",
				"name": "",
				"type": "address"
			}
		],
		"stateMutability": "view",
		"type": "function"
	},
	{
		"inputs": [
			{
				"internalType": "address",
				"name": "_candidate",
				"type": "address"
			}
		],
		"name": "ownerpropose",
		"outputs": [],
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"inputs": [
			{
				"internalType": "address",
				"name": "_candidate",
				"type": "address"
			}
		],
		"name": "ownerunpropose",
		"outputs": [],
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"inputs": [
			{
				"internalType": "address",
				"name": "_candidate",
				"type": "address"
			}
		],
		"name": "propose",
		"outputs": [],
		"stateMutability": "payable",
		"type": "function"
	},
	{
		"inputs": [
			{
				"internalType": "address",
				"name": "_candidate",
				"type": "address"
			}
		],
		"name": "resign",
		"outputs": [],
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"inputs": [
			{
				"internalType": "uint256[]",
				"name": "_rewards",
				"type": "uint256[]"
			}
		],
		"name": "rewardSetting",
		"outputs": [],
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"inputs": [
			{
				"internalType": "uint256",
				"name": "",
				"type": "uint256"
			}
		],
		"name": "rewards",
		"outputs": [
			{
				"internalType": "uint256",
				"name": "",
				"type": "uint256"
			}
		],
		"stateMutability": "view",
		"type": "function"
	},
	{
		"inputs": [
			{
				"internalType": "uint256",
				"name": "_minvalidatorcap",
				"type": "uint256"
			},
			{
				"internalType": "uint256",
				"name": "_maxvalidator",
				"type": "uint256"
			},
			{
				"internalType": "uint256",
				"name": "_validatorwdelay",
				"type": "uint256"
			},
			{
				"internalType": "uint256",
				"name": "_delegatorwdelay",
				"type": "uint256"
			}
		],
		"name": "settings",
		"outputs": [],
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"inputs": [
			{
				"internalType": "address",
				"name": "newOwner",
				"type": "address"
			}
		],
		"name": "transferOwnership",
		"outputs": [],
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"inputs": [
			{
				"internalType": "address",
				"name": "_candidate",
				"type": "address"
			},
			{
				"internalType": "uint256",
				"name": "_cap",
				"type": "uint256"
			}
		],
		"name": "undelegate",
		"outputs": [],
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"inputs": [],
		"name": "validatorCount",
		"outputs": [
			{
				"internalType": "uint256",
				"name": "",
				"type": "uint256"
			}
		],
		"stateMutability": "view",
		"type": "function"
	},
	{
		"inputs": [],
		"name": "validatorWithdrawDelay",
		"outputs": [
			{
				"internalType": "uint256",
				"name": "",
				"type": "uint256"
			}
		],
		"stateMutability": "view",
		"type": "function"
	},
	{
		"inputs": [
			{
				"internalType": "uint256",
				"name": "",
				"type": "uint256"
			}
		],
		"name": "validators",
		"outputs": [
			{
				"internalType": "address",
				"name": "",
				"type": "address"
			}
		],
		"stateMutability": "view",
		"type": "function"
	},
	{
		"inputs": [
			{
				"internalType": "uint256",
				"name": "_blockNumber",
				"type": "uint256"
			},
			{
				"internalType": "uint256",
				"name": "_index",
				"type": "uint256"
			}
		],
		"name": "withdraw",
		"outputs": [],
		"stateMutability": "nonpayable",
		"type": "function"
	}
]`
