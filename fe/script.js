
import { browser } from 'k6/experimental/browser';
import { check } from 'k6';

export const options = {
  scenarios: {
    ui: {
      executor: 'shared-iterations',
      options: {
        browser: {
          type: 'chromium',
        },
      },
    },
  },
  thresholds: {
    checks: ['rate==1.0'],
  },
};

export default async function () {
  const page = browser.newPage();

  try {
    await page.goto('http://localhost:8081/login');

    page.locator('input[name="login"]').type('frieren@gmail.com');
    page.locator('input[name="password"]').type('12341234');

    const submitButton = page.locator('input[type="submit"]');

    await Promise.all([page.waitForNavigation(), submitButton.click()]);

    check(page, {
      header: (p) => p.locator('span').textContent() == 'Hello, dinoman!',
    });
  } finally {
    page.close();
  }
}