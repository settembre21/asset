#!/bin/bash

pushd ../test-network > /dev/null

export PATH=${PWD}/../bin:$PATH
export FABRIC_CFG_PATH=$PWD/../config
export CORE_PEER_TLS_ENABLED=true
export CORE_PEER_LOCALMSPID="Org1MSP"
export CORE_PEER_TLS_ROOTCERT_FILE=${PWD}/organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt
export CORE_PEER_MSPCONFIGPATH=${PWD}/organizations/peerOrganizations/org1.example.com/users/Admin@org1.example.com/msp
export CORE_PEER_ADDRESS=localhost:7051

# ミラ・キータのVINコード
key='"JMYMIRAGINO200302"'

# 購入。locationが変わるだけ
mileage='""'
battery='""'
location='"Owner"'
args=${key},${year},${month},${mileage},${battery},${location}
echo "> ミラ・キータを購入しました。VINコードは[JMYMIRAGINO200302]です。"
echo "> Location:Owner"
date -u

peer chaincode invoke -o localhost:7050 --ordererTLSHostnameOverride orderer.example.com --tls --cafile ${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem -C mychannel -n asset --peerAddresses localhost:7051 --tlsRootCertFiles ${PWD}/organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt --peerAddresses localhost:9051 --tlsRootCertFiles ${PWD}/organizations/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/tls/ca.crt -c '{"function":"UpdateAsset","Args":['${args}']}'
sleep 5

# ディーラーへ修理依頼
mileage='"45678"'
battery='""'
location='"ＱＩＩＴＡ東京販売"'
args=${key},${mileage},${battery},${location}
echo "> 調子が悪いのでディーラーへ預けました。左ウィンカーが切れています。パワーウィンドウが動きません"
echo "> Location:ＱＩＩＴＡ東京販売"
date -u

peer chaincode invoke -o localhost:7050 --ordererTLSHostnameOverride orderer.example.com --tls --cafile ${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem -C mychannel -n asset --peerAddresses localhost:7051 --tlsRootCertFiles ${PWD}/organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt --peerAddresses localhost:9051 --tlsRootCertFiles ${PWD}/organizations/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/tls/ca.crt -c '{"function":"UpdateAsset","Args":['${args}']}'
sleep 5


# 工場へ転送
mileage='"45679"'
battery='""'
location='"ＱＩＩＴＡ埼玉工場"'
args=${key},${mileage},${battery},${location}
echo "> 工場へ運んで色々調べると連絡がありました。"
echo "> Location:ＱＩＩＴＡ埼玉工場"
date -u

peer chaincode invoke -o localhost:7050 --ordererTLSHostnameOverride orderer.example.com --tls --cafile ${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem -C mychannel -n asset --peerAddresses localhost:7051 --tlsRootCertFiles ${PWD}/organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt --peerAddresses localhost:9051 --tlsRootCertFiles ${PWD}/organizations/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/tls/ca.crt -c '{"function":"UpdateAsset","Args":['${args}']}'
sleep 5

# ディーラーへ戻る
mileage='"45680"'
battery='""'
location='"ＱＩＩＴＡ東京販売"'
args=${key},${mileage},${battery},${location}
echo "> 工場から戻ってきたと連絡がありました。いつ受け取りに行っても良いとのこと。"
echo "> Location:ＱＩＩＴＡ東京販売"
date -u

peer chaincode invoke -o localhost:7050 --ordererTLSHostnameOverride orderer.example.com --tls --cafile ${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem -C mychannel -n asset --peerAddresses localhost:7051 --tlsRootCertFiles ${PWD}/organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt --peerAddresses localhost:9051 --tlsRootCertFiles ${PWD}/organizations/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/tls/ca.crt -c '{"function":"UpdateAsset","Args":['${args}']}'
sleep 5

# オーナーへ戻る
mileage='"45680"'
battery='""'
location='"Owner"'
args=${key},${mileage},${battery},${location}
echo "> 愛しのミラ・キータが戻ってきました！"
echo "> Location:Owner"
date -u

peer chaincode invoke -o localhost:7050 --ordererTLSHostnameOverride orderer.example.com --tls --cafile ${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem -C mychannel -n asset --peerAddresses localhost:7051 --tlsRootCertFiles ${PWD}/organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt --peerAddresses localhost:9051 --tlsRootCertFiles ${PWD}/organizations/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/tls/ca.crt -c '{"function":"UpdateAsset","Args":['${args}']}'
sleep 5

# ディーラーへ廃車依頼
mileage='"67890"'
battery='""'
location='"ＱＩＩＴＡ東京販売"'
args=${key},${mileage},${battery},${location}
echo "> やっぱり調子が悪いのでディーラーに引き取ってもらいました。とても古いので廃車になるそうです。"
echo "> Location:ＱＩＩＴＡ東京販売"
date -u

peer chaincode invoke -o localhost:7050 --ordererTLSHostnameOverride orderer.example.com --tls --cafile ${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem -C mychannel -n asset --peerAddresses localhost:7051 --tlsRootCertFiles ${PWD}/organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt --peerAddresses localhost:9051 --tlsRootCertFiles ${PWD}/organizations/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/tls/ca.crt -c '{"function":"UpdateAsset","Args":['${args}']}'
sleep 5

# 工場へ転送
mileage='"67890"'
battery='""'
location='"ＱＩＩＴＡ千葉工場"'
args=${key},${mileage},${battery},${location}
echo "> 廃車にするため工場へ運ぶそうです"
echo "> Location:ＱＩＩＴＡ千葉工場"
date -u

peer chaincode invoke -o localhost:7050 --ordererTLSHostnameOverride orderer.example.com --tls --cafile ${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem -C mychannel -n asset --peerAddresses localhost:7051 --tlsRootCertFiles ${PWD}/organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt --peerAddresses localhost:9051 --tlsRootCertFiles ${PWD}/organizations/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/tls/ca.crt -c '{"function":"UpdateAsset","Args":['${args}']}'
sleep 5

# 廃車
args=${key}
echo "> さようならミラ・キータ。。。"
date -u

peer chaincode invoke -o localhost:7050 --ordererTLSHostnameOverride orderer.example.com --tls --cafile ${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem -C mychannel -n asset --peerAddresses localhost:7051 --tlsRootCertFiles ${PWD}/organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt --peerAddresses localhost:9051 --tlsRootCertFiles ${PWD}/organizations/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/tls/ca.crt -c '{"function":"DeleteAsset","Args":['${args}']}'
sleep 5

popd > /dev/null
