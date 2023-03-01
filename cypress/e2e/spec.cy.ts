describe('template-spec', () => {

  it('opens sign up page', () => {
    cy.visit('http://localhost:4200/signup')
    cy.get('.mat-mdc-button-touch-target')
    
  })

  it('opens login page', () => {
    cy.visit('http://localhost:4200/login')
  })
  
  it('opens profile page', () => {
    cy.visit('http://localhost:4200/profile')
  })
})