import React, { useState } from 'react';
import style from './transactions.module.css';
import { TransactionType } from '@/constants/response/resTransaction';
import { PropsConfigTransactions } from 'src/constants/types/props';
import {
  dropdownSortOptions,
  getSortLabel,
  getTypeName,
} from 'src/helpers/filterTransaction';

export const TransactionsFilter: React.FC<{
  transactionType: TransactionType[] | [];
  configTransaction: PropsConfigTransactions;
  setConfigTransaction: React.Dispatch<
    React.SetStateAction<PropsConfigTransactions>
  >;
}> = ({ transactionType, configTransaction, setConfigTransaction }) => {
  const [dropdownTypeVisible, setDropdownTypeVisible] = useState(false);
  const [dropdownSortVisible, setDropdownSortVisible] = useState(false);
  const [type, setType] = useState(
    getTypeName(transactionType, configTransaction),
  );
  const [sort, setSort] = useState(getSortLabel(configTransaction));

  function toggleDropdownType() {
    setDropdownTypeVisible(!dropdownTypeVisible);
  }

  function toggleDropdownSort() {
    setDropdownSortVisible(!dropdownSortVisible);
  }

  function chooseType(typeId: number, typeName: string) {
    setType(typeName);
    setConfigTransaction((prev) => ({
      ...prev,
      transactionType: typeId,
      page: 1,
    }));
    setDropdownTypeVisible(false);
  }

  function chooseSort(options: {
    sortBy: string;
    sortDir: string;
    label: string;
  }) {
    setSort(options.label);
    setConfigTransaction((prev) => ({
      ...prev,
      sortBy: options.sortBy,
      sortDir: options.sortDir,
      page: 1,
    }));
    setDropdownSortVisible(false);
  }

  return (
    <div className={style['tx-header-filters']}>
      <div className={style['tx-filter-group']}>
        <span>Type</span>
        <div className={style['tx-filter-container-all']}>
          <span
            role="presentation"
            onClick={toggleDropdownType}
            className={style['tx-filter-type']}
          >
            {type}
          </span>
          {dropdownTypeVisible && (
            <>
              <div>
                <div
                  role="presentation"
                  className={style['dropdown-type-sort-overlay']}
                  onClick={() => setDropdownTypeVisible(false)}
                ></div>
                <div className={style['tx-type-group']}>
                  {transactionType &&
                    transactionType.map((data, i) => (
                      <span
                        key={i + 1}
                        role="presentation"
                        onClick={() =>
                          chooseType(data?.type_id, data?.type_name)
                        }
                        className={style['tx-type-item']}
                      >
                        {data.type_name}
                      </span>
                    ))}
                </div>
              </div>
            </>
          )}
        </div>
      </div>
      <div className={style['tx-filter-group']}>
        <span>Sort</span>
        <div className={style['tx-filter-container-all']}>
          <span
            role="presentation"
            onClick={toggleDropdownSort}
            className={style['tx-filter-sort']}
          >
            {sort}
          </span>
          {dropdownSortVisible && (
            <>
              <div>
                <div
                  role="presentation"
                  className={style['dropdown-type-sort-overlay']}
                  onClick={() => setDropdownSortVisible(false)}
                ></div>
                <div className={style['tx-sort-group']}>
                  {dropdownSortOptions.map((option, i) => (
                    <span
                      key={i}
                      role="presentation"
                      onClick={() => chooseSort(option)}
                      className={style['tx-sort-item']}
                    >
                      {option.label}
                    </span>
                  ))}
                </div>
              </div>
            </>
          )}
        </div>
      </div>
    </div>
  );
};
