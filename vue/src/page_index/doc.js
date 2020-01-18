class Doc {
    id = "";
    title = "";
    subTitle = "";
    tags = [];
    attr = "";
    time = null;
    abstract = "";
    nextDoc = "";
    prefDoc = "";

    constructor(init = {}){
        this.id = init.id;
        this.title = init.title;
        this.subTitle = init.subTitle;
        this.tags = JSON.parse(JSON.stringify(init.tags));
        this.attr = init.attr;
        this.time = init.time;
        this.abstract = init.abstract;
        this.nextDoc = init.next;
        this.prefDoc = init.pref;
    }

    next(){
        return this.nextDoc
    }
    pref(){
        return this.prefDoc
    }
}

export default Doc