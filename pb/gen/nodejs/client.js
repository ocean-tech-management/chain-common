// cli := NewClient("3.39.236.151:8010", "testapi", []byte(TestPri), "Test")
const crypto = require("crypto");

module.exports.NewClient = function (host, pb_path, apikey, rsa_pri, name) {
  const protoLoader = require("@grpc/proto-loader");
  const grpcLibrary = require("@grpc/grpc-js");
  const packageDefinition = protoLoader.loadSync(pb_path);
  var proto = grpcLibrary.loadPackageDefinition(packageDefinition).proto;
  const client = new proto.ChainService(
    host,
    grpcLibrary.credentials.createInsecure()
  );
  const self = {};
  self.pri = rsa_pri;
  self.sign = function (data) {
    const sign = crypto.createSign("RSA-SHA256");
    sign.update(data);
    sign.end();
    const signature = sign.sign(self.pri);
    return signature.toString("base64");
  };
  self.apikey = apikey;
  self.CreateChainAccount = function (uid, cb) {
    const metadata = new grpcLibrary.Metadata();
    const nonce = crypto.randomUUID();
    metadata.set("nonce", nonce);
    metadata.set("api_key", apikey);
    metadata.set("name", name);
    const signdata =
      JSON.stringify({ api_key: self.apikey, uid: uid }) +
      "api_key=" +
      self.apikey +
      "&nonce=" +
      nonce;
    const signout = self.sign(signdata);
    metadata.set("sign", signout);
    client.CreateChainAccount(
      { apiKey: self.apikey, uid: uid },
      metadata,
      function (err, resp) {
        cb(err, resp);
      }
    );
  };
  self.SetNotifyUrl = function (url, cb) {
    const metadata = new grpcLibrary.Metadata();
    const nonce = crypto.randomUUID();
    metadata.set("nonce", nonce);
    metadata.set("api_key", apikey);
    metadata.set("name", name);
    const req = { url: url };
    const signdata =
      JSON.stringify(req) + "api_key=" + self.apikey + "&nonce=" + nonce;
    const signout = self.sign(signdata);
    metadata.set("sign", signout);
    client.SetNotifyUrl(req, metadata, function (err, resp) {
      cb(err, resp);
    });
  };
  return self;
};
