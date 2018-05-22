/**
 * @prettier
 */
// const host = "localhost";
// const tls = false;
const host = "192.168.86.56:8085";
const tls = false;

module.exports = {
    ApiHost: `${host}`,
    JwtStoreKey: "c4:jwt",
    UserSaveStoreKey: "c4:user",
    tls: tls,
};
