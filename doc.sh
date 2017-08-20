cd ./app && bee generate docs
cd ..
rm ./swagger/swagger.json
rm ./swagger/swagger.yml
cp ./app/swagger/* ./swagger
