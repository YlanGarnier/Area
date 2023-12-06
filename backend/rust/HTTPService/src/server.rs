use core::{core_service_client, ForwardActionReq};
use http_action::{
    http_service_action_server::{HttpServiceAction, HttpServiceActionServer},
    register_watcher,
};
use pmodel::Empty;
use tonic::{transport::Server, Request, Response, Status};


use futures::{stream::FuturesUnordered, StreamExt};
use prost::Message;
use std::sync::{Arc, Mutex};

pub mod http_action {
    tonic::include_proto!("http_action");
}

pub mod core {
    tonic::include_proto!("core");
}

pub mod pmodel {
    tonic::include_proto!("pmodel");
}

#[derive(Debug, Default, Clone)]
pub struct Watcher {
    id: u32,
    response_type: i32,
    url: String,
    expected: Vec<u16>,
}
#[derive(Debug, Default)]
pub struct HttpServiceActionService {
    watchers: Arc<Mutex<Vec<Watcher>>>,
}

impl HttpServiceActionService {
    pub fn new(watchers: Arc<Mutex<Vec<Watcher>>>) -> Self {
        Self { watchers }
    }
}

#[tonic::async_trait]
impl HttpServiceAction for HttpServiceActionService {
    async fn register_watcher(
        &self,
        request: Request<register_watcher::Request>,
    ) -> Result<Response<Empty>, Status> {
        let r = request.into_inner();
        println!("request begin: {:?}", r);
        let mut watchers = self.watchers.lock().unwrap();
        println!("request end: {:?}", r);
        watchers.push(Watcher {
            id: r.id,
            response_type: r.response_type,
            url: r.url,
            expected: r.expected.iter().map(|e| *e as u16).collect(),
        });

        Ok(Response::new(Empty {}))
    }
}

async fn watch(watcher: Watcher) {
    println!("watch");
    let resp = reqwest::get(&watcher.url).await;

    match resp {
        Ok(r) => {
            if watcher.expected.contains(&r.status().as_u16()) {
                return;
            };
            let mut client = core_service_client::CoreServiceClient::connect("http://127.0.0.1:9090").await.unwrap();
            let fmt_data = pmodel::format::GhIncidentReport {
                base: None,
                title: "Incident on localhost:3000".to_string(),
                content: "expected 200 but go another status".to_string(),
            };

            client
                .forward_action(Request::new(ForwardActionReq {
                    id: watcher.id,
                    r#type: watcher.response_type,
                    data: fmt_data.encode_to_vec(),
                }))
                .await;
                // .unwrap();
        }
        Err(e) => match e.status() {
            Some(s) => {
                if watcher.expected.contains(&s.as_u16()) {
                    return;
                };
                let mut client = core_service_client::CoreServiceClient::connect("http://127.0.0.1:9090").await.unwrap();
                let fmt_data = pmodel::format::GhIncidentReport {
                    base: None,
                    title: "Incident on localhost:3000".to_string(),
                    content: "expected 200 but go another status".to_string(),
                };

                client
                    .forward_action(Request::new(ForwardActionReq {
                        id: watcher.id,
                        r#type: watcher.response_type,
                        data: fmt_data.encode_to_vec(),
                    }))
                    .await;
                    // .unwrap();
            }
            None => {
                let mut client = core_service_client::CoreServiceClient::connect("http://127.0.0.1:9090").await.unwrap();
                let fmt_data = pmodel::format::GhIncidentReport {
                    base: None,
                    title: "Incident on localhost:3000".to_string(),
                    content: "expected 200 but go another status".to_string(),
                };

                client
                    .forward_action(Request::new(ForwardActionReq {
                        id: watcher.id,
                        r#type: watcher.response_type,
                        data: fmt_data.encode_to_vec(),
                    }))
                    .await;
                    // .unwrap();
            }
        },
    }
}

#[tokio::main]
async fn main() -> Result<(), Box<dyn std::error::Error>> {
    let watchers: Arc<Mutex<Vec<Watcher>>> = Arc::new(Mutex::new(Vec::new()));

    let watcherss = Arc::clone(&watchers);
    std::mem::drop(tokio::spawn(async move {
        loop {
            let futures = FuturesUnordered::new();
            {
                let wit = watcherss.lock().unwrap();
                for watcher in wit.iter() {
                    println!("watcher: {:?}", watcher);
                    futures.push(watch(watcher.clone()));
                }
            }

            // Using println! with futures doesn't provide useful information since
            // futures don't implement Debug in a meaningful way
            // println!("futures: {:?}", futures);
            futures.for_each(|_| futures::future::ready(())).await;

            // Use non-blocking sleep
            tokio::time::sleep(std::time::Duration::from_secs(5)).await;
        }
    }));

    let address = "127.0.0.1:9091".parse().unwrap();
    let service = HttpServiceActionService::new(watchers);

    Server::builder()
        .add_service(HttpServiceActionServer::new(service))
        .serve(address)
        .await?;
    println!("Server listening on {}", address);
    Ok(())
}

/*

for 1:
    if channel is not empty:
        get elements & add them to the list to poll
     poll the list

user:
1 lenny

services:
1 discord 1
2 telegram 1
3 slack 1


*/
