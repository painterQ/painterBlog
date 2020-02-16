if [ ! -d "./vue/node_modules" ]; then
  cd ./vue && npm install && cd ..
fi

cd ./vue && npm run build -config ./vue/vue.config.js && cd ..
cp -r ./vue/node_modules/tinymce/skins ./static/public/skins
go run ./