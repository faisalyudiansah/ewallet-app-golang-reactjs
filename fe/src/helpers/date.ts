export function getThisMonth(): string {
  const today = new Date();
  const month = today.toLocaleString('default', { month: 'long' });
  return month;
}

export function formatterDate(inputDate: string): string {
  const date = new Date(inputDate);
  return date.toLocaleDateString('id-ID', {
    day: 'numeric',
    month: 'long',
    year: 'numeric',
  });
}

export const FormatDate = (date: Date): string => {
  const options: Intl.DateTimeFormatOptions = {
    day: 'numeric',
    month: 'long',
    year: 'numeric',
  };
  return date.toLocaleDateString('id-ID', options);
};

export const formatDateToyyyymmdd = (date: Date): string => {
  const year = date.getFullYear();
  const month = String(date.getMonth() + 1).padStart(2, '0');
  const day = String(date.getDate()).padStart(2, '0');
  return `${year}-${month}-${day}`;
};

export const getThisWeek = (): {
  start_date: string;
  end_date: string;
} => {
  const now = new Date();
  const firstDayOfWeek = new Date(now.setDate(now.getDate() - now.getDay()));
  const lastDayOfWeek = new Date(now.setDate(now.getDate() - now.getDay() + 6));
  return {
    start_date: formatDateToyyyymmdd(firstDayOfWeek),
    end_date: formatDateToyyyymmdd(lastDayOfWeek),
  };
};

export const dataGetThisWeek = getThisWeek();
