POST   /v1/register              
JSON:       {
    "username": "dummyUname",
    "email": "dummy@Email.com",
    "password": ""dummypass",
    "confirm_password": "dummypass"
    }


POST   /v1/login
JSON:       {
    "username":"dummyUname",
    "password":"dummypass"
}

POST   /v1/Payment/
JSON:       {
    "receipt_username": "receiptUname",
    "transfer_amount": 20000
}


GET    /v1/wallet/

POST   /v1/wallet/topup
JSON:       {
    ""topup_amount: 20000
}