import http from "k6/http";

export const options = {
    iterations: 10,
};
const params = {
    headers: {
        'Authorization': 'Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MTI3OCwibmJmIjoxNzA0MTEwNDAwfQ.GP6RzCN4Z1bf_G3E_G2CsnA3QD4iq1YDLSvmVl7VySU',
    },
};


export default function () {
    const response = http.get("http://social-net:8080/user/search?first_name=И&second_name=Брагина", params);
}