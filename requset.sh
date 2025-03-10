curl -X POST "http://localhost:8080/transactions" \
     -H "Content-Type: application/json" \
     -d '{
       "user_id": 1,
       "tenor": 3,
       "contract_number": "TXN-20250308-0001",
       "otr": 25000000.00,
       "admin_fee": 500000.00,
       "installment_count": 3,
       "interest": 5.00,
       "asset_name": "Mobil X",
       "asset_type": "Mobil"
     }'
