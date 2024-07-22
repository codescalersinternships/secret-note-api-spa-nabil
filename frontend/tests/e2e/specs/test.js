// https://docs.cypress.io/api/table-of-contents

describe("End-to-End Test for Secret Note Application", () => {
  it("should sign up, sign in, create a note, and verify the note", () => {
    const uuid = () => Cypress._.random(0, 1e6);
    const id = uuid();
    const email = "test@gmail.com";

    // Sign Up
    cy.visit("http://localhost:8080/signup");
    cy.get('input[type="text"]').should("exist").type("nader");
    cy.get('input[type="email"]').should("exist").type(email);
    cy.get('input[type="password"]').should("exist").type("password");
    cy.get('button[type="submit"]').click();
    cy.url().should("include", "/login");

    // Sign In
    cy.get('input[type="email"]').should("exist").type(email);
    cy.get('input[type="password"]').type("password");
    cy.get('button[type="submit"]').click();
    cy.url().should("include", "/");

    // Create Note
    cy.visit("http://localhost:8080/create/");
    cy.get("textarea").should("exist").type("hello");
    cy.get('input[type="datetime-local"]')
      .should("exist")
      .type("2024-12-31T23:59");
    cy.get('input[type="number"]').should("exist").type("5");
    cy.get('button[type="submit"]').should("exist").click();
  });
});
