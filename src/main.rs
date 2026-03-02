use actix_web::*;
use tera::*;

mod handlers;

#[actix_web::main]
async fn main() -> std::io::Result<()> {
    return HttpServer::new(move || {
        App::new()
            .wrap(middleware::Compress::default())
            .app_data(web::Data::new(Tera::new("templates/*.html").unwrap()))

            .route("/",
                web::get().to(handlers::index))
            .route("/something",
                web::get().to(handlers::something))
    })
    .bind(("127.0.0.1", 6969))?
    .run()
    .await;
}
