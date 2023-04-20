describe('check login page', () => {
  it('passes', () => {
    cy.visit('http://localhost:4200')
  })
})

describe('check sign up page', () => {
  it('passes', () => {
    cy.visit('http://localhost:4200/signup')
  })
})

describe('check sidebar', () => {
  it('passes', () => {
    cy.visit('http://localhost:4200/sidebar')
  })
})

describe('check profile', () => {
  it('passes', () => {
    cy.visit('http://localhost:4200/profile')
  })
})

describe('check post', () => {
  it('passes', () => {
    cy.visit('http://localhost:4200/post')
  })
})