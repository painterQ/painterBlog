module.exports = {
  presets: [
    '@vue/app'
  ],

  //https://github.com/mAAdhaTTah/babel-plugin-prismjs
  "plugins": [
    ["prismjs", {
      "languages": ["js", "css", "markup", "go", "clike", "java"],
      "plugins": ["line-numbers","line-highlight"],
      "theme": "twilight",
      "css": true
    }]
  ]

}
