import { TransactionType } from 'src/constants/response/resTransaction';
import { PropsConfigTransactions } from 'src/constants/types/props';

export const dropdownSortOptions = [
  {
    sortBy: 'amount',
    sortDir: 'asc',
    label: 'Amount - Asc',
  },
  {
    sortBy: 'amount',
    sortDir: 'desc',
    label: 'Amount - Desc',
  },
  {
    sortBy: 'created_at',
    sortDir: 'asc',
    label: 'Date - Asc',
  },
  {
    sortBy: 'created_at',
    sortDir: 'desc',
    label: 'Date - Desc',
  },
];

export const getTypeName = (
  transactionType: TransactionType[] | [],
  configTransaction: PropsConfigTransactions,
): string => {
  const selectedType = transactionType.find(
    (type) => type.type_id === configTransaction.transactionType,
  );
  return selectedType ? selectedType.type_name : 'All';
};

export const getSortLabel = (
  configTransaction: PropsConfigTransactions,
): string => {
  const sortOption = dropdownSortOptions.find(
    (option) =>
      option.sortBy === configTransaction.sortBy &&
      option.sortDir === configTransaction.sortDir,
  );
  return sortOption ? sortOption.label : 'Amount - Asc';
};
