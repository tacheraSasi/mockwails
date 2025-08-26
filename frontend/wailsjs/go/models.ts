export namespace main {
	
	export class Server {
	    id: number;
	    name: string;
	    description: string;
	    endpoint: string;
	    method: string;
	    requestHeaders: string;
	    requestBody: string;
	    responseStatus: number;
	    responseHeaders: string;
	    responseBody: string;
	
	    static createFrom(source: any = {}) {
	        return new Server(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.name = source["name"];
	        this.description = source["description"];
	        this.endpoint = source["endpoint"];
	        this.method = source["method"];
	        this.requestHeaders = source["requestHeaders"];
	        this.requestBody = source["requestBody"];
	        this.responseStatus = source["responseStatus"];
	        this.responseHeaders = source["responseHeaders"];
	        this.responseBody = source["responseBody"];
	    }
	}

}

