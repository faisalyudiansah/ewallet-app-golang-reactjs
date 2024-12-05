import React from 'react';
import { fireEvent, render, screen } from '@testing-library/react';
import { HeroLeft } from '../Hero-left';
import { BrowserRouter } from 'react-router-dom';
const renderWithRouter = (ui: React.ReactElement, { route = '/' } = {}) => {
  window.history.pushState({}, 'Test page', route);
  return render(ui, { wrapper: BrowserRouter });
};

describe('test hero left on login page', () => {
  test('renders text', () => {
    renderWithRouter(<HeroLeft />);
    const textOnPage = [
      'Sign in to',
      'Sea Wallet',
      'If you don’t have an account register',
    ];
    textOnPage.forEach((text) => {
      const textPage = screen.getByText(text);
      expect(textPage).toBeInTheDocument();
    });
  });

  test('renders Link', () => {
    renderWithRouter(<HeroLeft />);
    const registerLink = screen.getByText('Register Here !');
    expect(registerLink).toHaveAttribute('href');
    fireEvent.click(registerLink);
    expect(window.location.pathname).toBe('/register');
  });
});
