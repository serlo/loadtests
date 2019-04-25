import http from "k6/http";
import { sleep } from "k6";

export let options = {
    vus: 10,
    duration: "30s"
    // stages: [
    //     { duration: "30s", target: 20 },
    //     { duration: "1m30s", target: 10  },
    //     { duration: "20s", target: 0 },
    // ]

};

export default function() {
    http.get("http://test.loadimpact.com");
    sleep(1);
};