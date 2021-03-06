import { assert } from 'chai';
import * as sinon from 'sinon';
import { ProtoJSONCompatible } from 'src/common';
import { EOFError, EOFMessage, MercuryWebSocket, IWebSocket } from './websocket';

class FakeMessage implements ProtoJSONCompatible {
	id?: string;
	num?: number;
	worked?: boolean;
	ToProtoJSON(): Object {
		return {
			id: this.id,
			num: this.num?.toString(),
			worked: this.worked
		}
	}
}

class FakeResponse {
	resultData?: string;
}

async function ParseFakeData(res: any): Promise<FakeResponse> {
	let parsed = JSON.parse(res);
	let resT: FakeResponse = {
		resultData: parsed.resultData
	};
	return resT;
}

function fakeWsConstructor(): IWebSocket {
	throw new Error("Unimplemented")
}

function msgEventConstructor(type: string): MessageEvent {
	return {} as any;
}

let fileSandbox = sinon.createSandbox();
before(async () => {
	global.WebSocket = fakeWsConstructor as any;
	global.MessageEvent = msgEventConstructor as any;
	fileSandbox.stub(global, "WebSocket").callsFake(fakeWsConstructor);
	fileSandbox.stub(global, "MessageEvent").callsFake(msgEventConstructor);
})
after(async () => {
	fileSandbox.restore();
})

describe("Websocket", () => {
	describe("Basic operations", () => {
		let sandbox: sinon.SinonSandbox;
		beforeEach(async () => {
			sandbox = sinon.createSandbox();
		})
		afterEach(async () => {
			sandbox.restore();
		})
		it("Init succeeds", async () => {
			let fake = new FakeWebsocket();
			let evStub = sandbox.stub(fake, "addEventListener");
			let evsCalled = [false, false, false, false];
			evStub.withArgs("close", sinon.match(() => true)).callsFake(() => {
				evsCalled[0] = true;
			})
			evStub.withArgs("open", sinon.match(() => true)).callsFake(() => {
				evsCalled[1] = true;
			})
			evStub.withArgs("message", sinon.match(() => true)).callsFake(() => {
				evsCalled[2] = true;
			})
			evStub.withArgs("error", sinon.match(() => true)).callsFake(() => {
				evsCalled[3] = true;
			})
			let ws = new MercuryWebSocket<FakeMessage, FakeResponse>("not a real URL", ParseFakeData, "TestWebsocket", console, () => fake);
			await ws.init();
			assert.deepEqual(evsCalled, [true, true, true, true]);
		});
		it("Send works", async () => {
			let mockedWs = await makeMockedWebsocket(sandbox);
			let ws = mockedWs.ws;
			let fake = mockedWs.fake;
			sandbox.stub(fake, "send").callsFake(data => {
				assert.equal(data, `{"id":"abcd"}`)
			})
			mockedWs.open({} as any);
			let msg = new FakeMessage();
			msg.id = "abcd";
			await ws.Send(msg);
		});
		it("Recv works", async () => {
			let mockedWs = await makeMockedWebsocket(sandbox);
			let ws = mockedWs.ws;
			mockedWs.open({} as any);
			mockedWs.message({ data: `{"resultData":"test"}` } as any);
			let res = await ws.Recv();
			let expectedRes = new FakeResponse();
			expectedRes.resultData = "test";
			assert.deepEqual(res, expectedRes);
		});
		it("Send and Recv work together", async () => {
			let mockedWs = await makeMockedWebsocket(sandbox);
			let ws = mockedWs.ws;
			let fake = mockedWs.fake;
			sandbox.stub(fake, "send").callsFake(data => {
				assert.equal(data, `{"id":"abcd"}`)
			})
			mockedWs.open({} as any);
			mockedWs.message({ data: `{"resultData":"test"}` } as any);
			let msg = new FakeMessage();
			msg.id = "abcd";
			await ws.Send(msg);
			let res = await ws.Recv();
			let expectedRes = new FakeResponse();
			expectedRes.resultData = "test";
			assert.deepEqual(res, expectedRes);
		});
		it("EOF is handled", async () => {
			let mockedWs = await makeMockedWebsocket(sandbox);
			let ws = mockedWs.ws;
			mockedWs.open({} as any);
			mockedWs.message({ data: EOFMessage } as any);
			try {
				await ws.Recv();
				assert.fail("should not have completed Recv")
			} catch (err) {
				if (!(err instanceof EOFError)) {
					assert.fail("must be EOF error")
				}
			}
		});
	})


})

class FakeWebsocket implements IWebSocket {
	send(data: string | ArrayBuffer | SharedArrayBuffer | Blob | ArrayBufferView): void {
		throw new Error("no websocket")
	}
	close(code?: number | undefined, reason?: string | undefined): void {
		throw new Error("no websocket")
	}
	addEventListener<K extends keyof WebSocketEventMap>(type: K, listener: (this: WebSocket, ev: WebSocketEventMap[K]) => any, options?: boolean | AddEventListenerOptions): void
	addEventListener(type: string, listener: EventListenerOrEventListenerObject, options?: boolean | AddEventListenerOptions): void {
		throw new Error("no websocket")
	}
}

interface mockedWebsocket {
	ws: MercuryWebSocket<FakeMessage, FakeResponse>;
	fake: IWebSocket;
	close(ev: CloseEvent): void;
	open(ev: Event): void;
	message(ev: MessageEvent): void;
	error(ev: Event): void;
}

async function makeMockedWebsocket(sandbox: sinon.SinonSandbox): Promise<mockedWebsocket> {
	let fake = new FakeWebsocket();
	let evStub = sandbox.stub(fake, "addEventListener");
	let done: ((a: any) => any)[] = [];
	let wait: [Promise<(ev: CloseEvent) => void>, Promise<(ev: Event) => void>, Promise<(ev: MessageEvent) => void>, Promise<(ev: Event) => void>] = [
		new Promise(resolve => {
			done.push(resolve);
		}),
		new Promise(resolve => {
			done.push(resolve);
		}),
		new Promise(resolve => {
			done.push(resolve);
		}),
		new Promise(resolve => {
			done.push(resolve);
		}),
	];
	evStub.withArgs("close", sinon.match(() => true)).callsFake((type, listener) => {
		done[0](listener);
	})
	evStub.withArgs("open", sinon.match(() => true)).callsFake((type, listener) => {
		done[1](listener);
	})
	evStub.withArgs("message", sinon.match(() => true)).callsFake((type, listener) => {
		done[2](listener);
	})
	evStub.withArgs("error", sinon.match(() => true)).callsFake((type, listener) => {
		done[3](listener);
	})
	let ws = new MercuryWebSocket<FakeMessage, FakeResponse>("not a real URL", ParseFakeData, "TestWebsocket", console, () => fake);
	await ws.init();
	let [close, open, message, error] = await Promise.all(wait);
	return {
		ws,
		fake,
		close,
		open,
		message,
		error
	}
}