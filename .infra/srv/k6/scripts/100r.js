import http from 'k6/http';
import { check } from 'k6';

const params = {
    headers: {
        'Authorization': 'Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MTI3OCwibmJmIjoxNzA0MTEwNDAwfQ.GP6RzCN4Z1bf_G3E_G2CsnA3QD4iq1YDLSvmVl7VySU',
    },
};


export default function () {
    const countRequest = 100
    let request = []

    for (let req = 0; req < countRequest; req++) {
        request.push({
            method: 'GET',
            url: 'http://social-net:8080/user/search?first_name=И&second_name=Брагина',
            params: params
        });
    }



    const responses = http.batch(request);

    // httpbin.test.k6.io should return our POST data in the response body, so
    // we check the third response object to see that the POST worked.
    // check(responses[2], {
    //     'form data OK': (res) => JSON.parse(res.body)['form']['hello'] == 'world!',
    // });
}
