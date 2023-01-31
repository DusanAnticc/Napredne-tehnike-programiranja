#![feature(proc_macro_hygiene, decl_macro)]

#[macro_use] extern crate rocket;

use lettre::{
    transport::smtp::authentication::Credentials, AsyncSmtpTransport, AsyncTransport, Message,
    Tokio1Executor,
};


#[tokio::main]
#[get("/")]
async fn index() -> Result<(), Box<dyn std::error::Error>> {
    let smtp_credentials =
    Credentials::new("dusan.antic177@gmail.com".to_string(), "blabla".to_string());

let mailer = AsyncSmtpTransport::<Tokio1Executor>::relay("smtp.email.com")?
    .credentials(smtp_credentials)
    .build();

let from = "dusan.antic177@gmail.com";
let to = "dusan.antic177@gmail.com";
let subject = "Sending email with Rust";
let body = "you have new notification".to_string();

send_email_smtp(&mailer, from, to, subject, body).await
}


fn main() {    
    rocket::ignite().mount("/", routes![index]).launch();
}

// Email sending function
async fn send_email_smtp(
    mailer: &AsyncSmtpTransport<Tokio1Executor>,
    from: &str,
    to: &str,
    subject: &str,
    body: String,
) -> Result<(), Box<dyn std::error::Error>> {
    let email = Message::builder()
        .from(from.parse()?)
        .to(to.parse()?)
        .subject(subject)
        .body(body.to_string())?;

    mailer.send(email).await?;

    Ok(())
}
