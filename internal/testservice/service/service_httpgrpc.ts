/**
 * Code generated by httpgrpc. DO NOT EDIT.
 * versions:
 * 	httpgrpc v0.8.0
 * 	protoc   v3.10.1
 * source: service.proto
 */

import * as packages from "__packages__";
import * as httpgrpc_packages from "__packages__/httpgrpc";
import * as httpgrpc from "@llkennedy/httpgrpc";

/** A service client */
export class ExposedAppClient extends httpgrpc.Client {
	public async Random(req: httpgrpc_packages.service.RandomRequest): Promise<httpgrpc_packages.service.RandomResponse> {
		return this.SendUnary<httpgrpc_packages.service.RandomRequest, httpgrpc_packages.service.RandomResponse>("/Random", httpgrpc.HTTPMethod.GET, req, httpgrpc_packages.service.RandomResponse.Parse);
	}
	public async UploadPhoto(req: httpgrpc_packages.service.UploadPhotoRequest): Promise<httpgrpc_packages.service.UploadPhotoResponse> {
		return this.SendUnary<httpgrpc_packages.service.UploadPhotoRequest, httpgrpc_packages.service.UploadPhotoResponse>("/UploadPhoto", httpgrpc.HTTPMethod.POST, req, httpgrpc_packages.service.UploadPhotoResponse.Parse);
	}
	public async Feed(): Promise<httpgrpc.IClientStream<httpgrpc_packages.service.FeedData, httpgrpc_packages.service.FeedResponse>> {
		return this.StartClientStream<httpgrpc_packages.service.FeedData, httpgrpc_packages.service.FeedResponse>("/Feed", httpgrpc_packages.service.FeedResponse.Parse);
	}
	public async Broadcast(req: httpgrpc_packages.service.BroadcastRequest): Promise<httpgrpc.IServerStream<httpgrpc_packages.service.BroadcastData>> {
		return this.StartServerStream<httpgrpc_packages.service.BroadcastRequest, httpgrpc_packages.service.BroadcastData>("/Broadcast", req, httpgrpc_packages.service.BroadcastData.Parse);
	}
	public async ConvertToString(): Promise<httpgrpc.IDualStream<httpgrpc_packages.service.ConvertInput, httpgrpc_packages.service.ConvertOutput>> {
		return this.StartDualStream<httpgrpc_packages.service.ConvertInput, httpgrpc_packages.service.ConvertOutput>("/ConvertToString", httpgrpc_packages.service.ConvertOutput.Parse);
	}
}

/** A message */
export class FeedData extends packages.service.FeedData implements httpgrpc.ProtoJSONCompatible {
	public ToProtoJSON(): Object {
		return {
			id: httpgrpc.ToProtoJSON.String(this.id),
			dataType: httpgrpc.ToProtoJSON.StringNumber(this.dataType),
			rawData: httpgrpc.ToProtoJSON.Bytes(this.rawData),
		};
	}
	public static async Parse(data: any): Promise<FeedData> {
		let objData: Object = httpgrpc.AnyToObject(data);
		let res = new FeedData();
		res.id = await httpgrpc.Parse.String(objData, "id", "id");
		res.dataType = await httpgrpc.Parse.Number(objData, "dataType", "data_type");
		res.rawData = await httpgrpc.Parse.Bytes(objData, "rawData", "raw_data");
		return res;
	}
}

/** A message */
export class FeedResponse extends packages.service.FeedResponse implements httpgrpc.ProtoJSONCompatible {
	public ToProtoJSON(): Object {
		return {
			received: httpgrpc.ToProtoJSON.Number(this.received),
		};
	}
	public static async Parse(data: any): Promise<FeedResponse> {
		let objData: Object = httpgrpc.AnyToObject(data);
		let res = new FeedResponse();
		res.received = await httpgrpc.Parse.Number(objData, "received", "received");
		return res;
	}
}

/** A message */
export class BroadcastRequest extends packages.service.BroadcastRequest implements httpgrpc.ProtoJSONCompatible {
	public ToProtoJSON(): Object {
		return {
			id: httpgrpc.ToProtoJSON.String(this.id),
		};
	}
	public static async Parse(data: any): Promise<BroadcastRequest> {
		let objData: Object = httpgrpc.AnyToObject(data);
		let res = new BroadcastRequest();
		res.id = await httpgrpc.Parse.String(objData, "id", "id");
		return res;
	}
}

/** A message */
export class BroadcastData extends packages.service.BroadcastData implements httpgrpc.ProtoJSONCompatible {
	public ToProtoJSON(): Object {
		return {
			rawData: httpgrpc.ToProtoJSON.Bytes(this.rawData),
		};
	}
	public static async Parse(data: any): Promise<BroadcastData> {
		let objData: Object = httpgrpc.AnyToObject(data);
		let res = new BroadcastData();
		res.rawData = await httpgrpc.Parse.Bytes(objData, "rawData", "raw_data");
		return res;
	}
}

/** A message */
export class ConvertInput extends packages.service.ConvertInput implements httpgrpc.ProtoJSONCompatible {
	public ToProtoJSON(): Object {
		return {
			rawData: httpgrpc.ToProtoJSON.Bytes(this.rawData),
		};
	}
	public static async Parse(data: any): Promise<ConvertInput> {
		let objData: Object = httpgrpc.AnyToObject(data);
		let res = new ConvertInput();
		res.rawData = await httpgrpc.Parse.Bytes(objData, "rawData", "raw_data");
		return res;
	}
}

/** A message */
export class ConvertOutput extends packages.service.ConvertOutput implements httpgrpc.ProtoJSONCompatible {
	public ToProtoJSON(): Object {
		return {
			convertedData: httpgrpc.ToProtoJSON.String(this.convertedData),
		};
	}
	public static async Parse(data: any): Promise<ConvertOutput> {
		let objData: Object = httpgrpc.AnyToObject(data);
		let res = new ConvertOutput();
		res.convertedData = await httpgrpc.Parse.String(objData, "convertedData", "converted_data");
		return res;
	}
}

/** A message */
export class FibonacciRequest extends packages.service.FibonacciRequest implements httpgrpc.ProtoJSONCompatible {
	public ToProtoJSON(): Object {
		return {
			n: httpgrpc.ToProtoJSON.StringNumber(this.n),
		};
	}
	public static async Parse(data: any): Promise<FibonacciRequest> {
		let objData: Object = httpgrpc.AnyToObject(data);
		let res = new FibonacciRequest();
		res.n = await httpgrpc.Parse.Number(objData, "n", "n");
		return res;
	}
}

/** A message */
export class FibonacciResponse extends packages.service.FibonacciResponse implements httpgrpc.ProtoJSONCompatible {
	public ToProtoJSON(): Object {
		return {
			number: httpgrpc.ToProtoJSON.StringNumber(this.number),
		};
	}
	public static async Parse(data: any): Promise<FibonacciResponse> {
		let objData: Object = httpgrpc.AnyToObject(data);
		let res = new FibonacciResponse();
		res.number = await httpgrpc.Parse.Number(objData, "number", "number");
		return res;
	}
}

/** A message */
export class RandomRequest extends packages.service.RandomRequest implements httpgrpc.ProtoJSONCompatible {
	public ToProtoJSON(): Object {
		return {
			lowerBound: httpgrpc.ToProtoJSON.StringNumber(this.lowerBound),
			upperBound: httpgrpc.ToProtoJSON.StringNumber(this.upperBound),
		};
	}
	public static async Parse(data: any): Promise<RandomRequest> {
		let objData: Object = httpgrpc.AnyToObject(data);
		let res = new RandomRequest();
		res.lowerBound = await httpgrpc.Parse.Number(objData, "lowerBound", "lower_bound");
		res.upperBound = await httpgrpc.Parse.Number(objData, "upperBound", "upper_bound");
		return res;
	}
}

/** A message */
export class RandomResponse extends packages.service.RandomResponse implements httpgrpc.ProtoJSONCompatible {
	public ToProtoJSON(): Object {
		return {
			number: httpgrpc.ToProtoJSON.StringNumber(this.number),
		};
	}
	public static async Parse(data: any): Promise<RandomResponse> {
		let objData: Object = httpgrpc.AnyToObject(data);
		let res = new RandomResponse();
		res.number = await httpgrpc.Parse.Number(objData, "number", "number");
		return res;
	}
}

/** A message */
export class UploadPhotoRequest extends packages.service.UploadPhotoRequest implements httpgrpc.ProtoJSONCompatible {
	public ToProtoJSON(): Object {
		return {
			data: httpgrpc.ToProtoJSON.Bytes(this.data),
		};
	}
	public static async Parse(data: any): Promise<UploadPhotoRequest> {
		let objData: Object = httpgrpc.AnyToObject(data);
		let res = new UploadPhotoRequest();
		res.data = await httpgrpc.Parse.Bytes(objData, "data", "data");
		return res;
	}
}

/** A message */
export class UploadPhotoResponse extends packages.service.UploadPhotoResponse implements httpgrpc.ProtoJSONCompatible {
	public ToProtoJSON(): Object {
		return {
			uuid: httpgrpc.ToProtoJSON.String(this.uuid),
		};
	}
	public static async Parse(data: any): Promise<UploadPhotoResponse> {
		let objData: Object = httpgrpc.AnyToObject(data);
		let res = new UploadPhotoResponse();
		res.uuid = await httpgrpc.Parse.String(objData, "uuid", "uuid");
		return res;
	}
}

