# LOGS of my commands

## Transactions (Transfers)


Sending the TX
```
➜  ~/QEVM/go-cellframe-cli git:(master) ✗ go run src/main.go
2025/01/23 15:26:52 Executing transaction creation command: /opt/cellframe-node/bin/cellframe-node-cli [tx_create -net Backbone -chain main -value 100000000000000000 -token CELL -to_addr Rj7J7MiX2bWy8sNyb2Zn74XqguhasMmiVvmM6amcQVyAmCTuqfutnFa8wfe4ts93cuDdEdUDQ59fCbbqBMisvzRySWCv6Hp6sDS1kbjz -from_wallet exchange -fee 50000000000000000]
Successfully created transaction:
transfer: Ok
        hash: 0x7DA343D900A7023699591B468D9D82162E7B3551AB521CE912FF8A4B1B9F1FCE
```


Checking the output
```
> tx_history -addr Rj7J7MiX2bWy8sNyZDhn63wrZbC71FjnEMJEEC9vYfMTdpkdo17iPH5A2ecsuCbs9WPt4kJ3wAEPi1upocqVBsV4jstkHCtX8g8L1Jyr -net Backbone
    
        address: Rj7J7MiX2bWy8sNyZDhn63wrZbC71FjnEMJEEC9vYfMTdpkdo17iPH5A2ecsuCbs9WPt4kJ3wAEPi1upocqVBsV4jstkHCtX8g8L1Jyr

    
        limit: unlimit

    
        status: ACCEPTED
        hash: 0x7DA343D900A7023699591B468D9D82162E7B3551AB521CE912FF8A4B1B9F1FCE
        atom_hash: 0x62877708243B4825EA68EC52F0023100D5E82B5E14CF968901F9D53B172410EC
        ret_code: 0
        ret_code_str: No error
        action: regular
        service: transfer
        batching: false
        tx_created: Thu, 23 Jan 2025 15:26:52 +0100
        data: 
            
                tx_type: send
                send_coins: 0.1
                send_datoshi: 100000000000000000
                token: CELL
                destination_address: Rj7J7MiX2bWy8sNyb2Zn74XqguhasMmiVvmM6amcQVyAmCTuqfutnFa8wfe4ts93cuDdEdUDQ59fCbbqBMisvzRySWCv6Hp6sDS1kbjz

            
                tx_type: send
                send_coins: 0.05
                send_datoshi: 50000000000000000
                token: CELL
                destination_address: DAP_CHAIN_TX_OUT_COND_SUBTYPE_FEE



    
        status: ACCEPTED
        hash: 0xD2B8E8CC6BFE969F0776BDE397B7DAD13BD4B08C10351A426BF992A0BEE1ED35
        atom_hash: 0x40D299E615E877E336C3FF4269171231CD0C6B5A322F68C70A933FF73F4B1CED
        ret_code: 0
        ret_code_str: No error
        action: regular
        service: transfer
        batching: false
        tx_created: Thu, 23 Jan 2025 15:11:56 +0100
        data: 
            
                tx_type: recv
                recv_coins: 1.0
                recv_datoshi: 1000000000000000000
                token: CELL
                source_address: Rj7J7MiX2bWy8sNyb2Zn74XqguhasMmiVvmM6amcQVyAmCTuqfutnFa8wfe4ts93cuDdEdUDQ59fCbbqBMisvzRySWCv6Hp6sDS1kbjz




Print 2 transactions in network Backbone chain main. 
Of which 2 were accepted into the ledger and 0 were rejected.
```




## Exchanges

Purchasing QEVM - randomly found order
```
➜  ~/QEVM/go-cellframe-cli git:(master) ✗ go run src/main.go                                              
2025/01/23 15:36:55 Executing transaction creation command: /opt/cellframe-node/bin/cellframe-node-cli [srv_xchange purchase -order 0x3731105E08B5633B546A2D9FDC1C3A02FC7022EF6221AF6D4F7A6B80DB22C965 -net Backbone -w exchange -value 10000000000000000 -fee 50000000000000000]
Successfully purchased order:
Exchange transaction has done. tx hash: 0x3C696E1D9D93BEE9A3C05DDF04884F75785C7EB9D580CC31FC3FC37A4DCC0DDA

```

Checking results & output afterwards:
```
> wallet info -w exchange -net Backbone
    
        
            sign: correct
            wallet: exchange
            addr: Rj7J7MiX2bWy8sNyZDhn63wrZbC71FjnEMJEEC9vYfMTdpkdo17iPH5A2ecsuCbs9WPt4kJ3wAEPi1upocqVBsV4jstkHCtX8g8L1Jyr
            network: Backbone
            signs: sig_dil
            tokens: 
                
                    balance: 
                    coins: 0.79
                    datoshi: 790000000000000000
                    token: 
                        ticker: CELL
                        description: 


                
                    balance: 
                    coins: 0.000015
                    datoshi: 15000000000000
                    token: 
                        ticker: QEVM
                        description: 





> 
```
