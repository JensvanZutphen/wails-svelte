export namespace models {
	
	export class PokemonSpecies {
	    description: string;
	
	    static createFrom(source: any = {}) {
	        return new PokemonSpecies(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.description = source["description"];
	    }
	}
	export class PokemonAbility {
	    name: string;
	    description: string;
	    is_hidden: boolean;
	
	    static createFrom(source: any = {}) {
	        return new PokemonAbility(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.name = source["name"];
	        this.description = source["description"];
	        this.is_hidden = source["is_hidden"];
	    }
	}
	export class PokemonStat {
	    name: string;
	    base_stat: number;
	
	    static createFrom(source: any = {}) {
	        return new PokemonStat(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.name = source["name"];
	        this.base_stat = source["base_stat"];
	    }
	}
	export class Pokemon {
	    id: number;
	    name: string;
	    image_url: string;
	    types: string[];
	    stats: PokemonStat[];
	    abilities: PokemonAbility[];
	    height: number;
	    weight: number;
	    species: PokemonSpecies;
	
	    static createFrom(source: any = {}) {
	        return new Pokemon(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.name = source["name"];
	        this.image_url = source["image_url"];
	        this.types = source["types"];
	        this.stats = this.convertValues(source["stats"], PokemonStat);
	        this.abilities = this.convertValues(source["abilities"], PokemonAbility);
	        this.height = source["height"];
	        this.weight = source["weight"];
	        this.species = this.convertValues(source["species"], PokemonSpecies);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	
	

}

