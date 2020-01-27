class Doc {
    id = "";
    title = "";
    subTitle = "";
    tags = [];
    attr = "";
    time = null;
    abstract = "";

    constructor(init = {}){
        this.id = init.id;
        this.title = init.title;
        this.subTitle = init.subTitle;
        this.tags = JSON.parse(JSON.stringify(init.tags));
        this.attr = init.attr;
        this.time = init.time;
        this.abstract = init.abstract;
    }
}

export default Doc