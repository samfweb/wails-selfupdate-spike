export namespace models {
	
	export class Subscription {
	    id: number;
	    // Go type: time
	    createdAt: any;
	    // Go type: time
	    updatedAt: any;
	    // Go type: gorm
	    deletedAt?: any;
	    connectionId: number;
	    qos: number;
	    topic: string;
	    protoDescriptor: string;
	
	    static createFrom(source: any = {}) {
	        return new Subscription(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.createdAt = this.convertValues(source["createdAt"], null);
	        this.updatedAt = this.convertValues(source["updatedAt"], null);
	        this.deletedAt = this.convertValues(source["deletedAt"], null);
	        this.connectionId = source["connectionId"];
	        this.qos = source["qos"];
	        this.topic = source["topic"];
	        this.protoDescriptor = source["protoDescriptor"];
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
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
	export class Connection {
	    id: number;
	    // Go type: time
	    createdAt: any;
	    // Go type: time
	    updatedAt: any;
	    // Go type: gorm
	    deletedAt?: any;
	    name: string;
	    protocol: string;
	    host: string;
	    port: number;
	    // Go type: JsonNullString
	    username: any;
	    // Go type: JsonNullString
	    password: any;
	    isProtoEnabled: boolean;
	    // Go type: JsonNullString
	    protoRegDir: any;
	    isCertsEnabled: boolean;
	    // Go type: JsonNullString
	    certCa: any;
	    // Go type: JsonNullString
	    certClient: any;
	    // Go type: JsonNullString
	    certClientKey: any;
	    subscriptions: Subscription[];
	
	    static createFrom(source: any = {}) {
	        return new Connection(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.createdAt = this.convertValues(source["createdAt"], null);
	        this.updatedAt = this.convertValues(source["updatedAt"], null);
	        this.deletedAt = this.convertValues(source["deletedAt"], null);
	        this.name = source["name"];
	        this.protocol = source["protocol"];
	        this.host = source["host"];
	        this.port = source["port"];
	        this.username = this.convertValues(source["username"], null);
	        this.password = this.convertValues(source["password"], null);
	        this.isProtoEnabled = source["isProtoEnabled"];
	        this.protoRegDir = this.convertValues(source["protoRegDir"], null);
	        this.isCertsEnabled = source["isCertsEnabled"];
	        this.certCa = this.convertValues(source["certCa"], null);
	        this.certClient = this.convertValues(source["certClient"], null);
	        this.certClientKey = this.convertValues(source["certClientKey"], null);
	        this.subscriptions = this.convertValues(source["subscriptions"], Subscription);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
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

