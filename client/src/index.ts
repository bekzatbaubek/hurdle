console.info("Loaded Hurdle");

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
  });
}
