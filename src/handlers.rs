use actix_web::http::header::{ContentType};
use actix_web::{HttpResponse, web};
use tera::*;

pub async fn index(
    t: web::Data<Tera>
) -> HttpResponse {
    let datastar_js = include_str!("../templates/static/datastar.js");

    let mut context = Context::new();
    context.insert("datastar_js", datastar_js);

    let page = t.render(
        "index.html", &context);

    return HttpResponse::Ok()
        .content_type("text/html")
        .body(page.unwrap());
}

pub async fn something() -> HttpResponse {
    let response = r#"
        <div id="hal">Something</div>
    "#.to_string();

    return HttpResponse::Ok()
        .insert_header(ContentType::html())
        .body(response);
}
