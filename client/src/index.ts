const baseUrl = "http://localhost:8080";

console.info("Loaded Hurdle");

interface Guess {
  round: number;
  guess: string;
  hint: string;
}

interface GameSession {
  current_round: number;
  guesses: Guess[];
}

async function createOrLoadSession() {
  let response = await fetch(`${baseUrl}/api/sessions`, {
    credentials: "include",
  });
  let session: GameSession = await response.json();
  console.log("Session:", session);
}

createOrLoadSession();

let hurdleForm = document.querySelector("#hurdle-form");

if (!hurdleForm) {
  console.error("Could not find hurdle form");
} else {
  hurdleForm.addEventListener("submit", (event) => {
    event.preventDefault();
    console.log("Hurdle form submitted");

    let formData = new FormData(hurdleForm as HTMLFormElement);

    let guess = formData.get("hurdle-guess") as string;
    console.log("Guess:", guess);

    fetch(`${baseUrl}/api/guesses`, {
      method: "POST",
      credentials: "include",
      body: JSON.stringify({ guess: guess }),
    }).then((response) => {
      response.json().then((data) => {
        console.log("Response data:", data);
      });
    });
  });
}
