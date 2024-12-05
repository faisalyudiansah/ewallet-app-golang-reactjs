export function CurrencyFormatterIDR(amount: number): string {
  let money = amount;
  if (amount < 0) {
    money = amount * -1;
  }
  return money.toLocaleString('id-ID').replace(/\./g, ',');
}

export function CurrencyFormatterIDRInput(amount: number): string {
  return amount.toLocaleString('id-ID');
}
