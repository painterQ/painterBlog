class Doc {
    id = "";
    title = "";
    subTitle = "";
    tags = [];
    attr = "";
    time = null;
    abstract = "";
    /*
    * {
      "id": "/second",
      "title": "第二篇",
      "subTitle": "blog",
      "tags": [
        "blog",
        "second"
      ],
      "attr": 0,
      "time": 1580263731,
      "abstract": "这是我的第二篇blog，希望从此没有bug。"
    }
    * */
    constructor(init = {}){
        this.id = init.id;
        this.title = init.title;
        this.subTitle = init.subTitle;
        this.tags = JSON.parse(JSON.stringify(init.tags));
        this.attr = init.attr;
        this.time = Number(init.time) * 1000;
        this.abstract = init.abstract;
    }
}

export default Doc