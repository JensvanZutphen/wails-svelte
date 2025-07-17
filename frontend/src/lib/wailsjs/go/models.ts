export namespace main {
	
	export class Pokemon {
	    id: number;
	    name: string;
	    image_url: string;
	    types: string[];
	
	    static createFrom(source: any = {}) {
	        return new Pokemon(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.name = source["name"];
	        this.image_url = source["image_url"];
	        this.types = source["types"];
	    }
	}

}

