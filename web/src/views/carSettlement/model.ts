import { h, ref } from 'vue';
import { NTag } from 'naive-ui';
import { cloneDeep } from 'lodash-es';
import { Dicts } from '@/api/dict/dict';
import { isNullObject } from '@/utils/is';
import { Option, getOptionLabel, getOptionTag } from '@/utils/hotgo';

export class State {
  public id = 0; // id
  public name = '';// 名称
  public type = 1; // 结算方式
  public cycle = 1; // 结算周期
  public rate = 0; // 结算比例
  public costArr = null; // 结算成本控制
  public cost = null; // 结算成本控制
  public isAuditWithdraw = 1; // 是否提现审核
  public status = 1; // 状态
  public createAt = ''; // 创建时间
  public updateAt = ''; // 更新时间
  public deletedAt = ''; // deleted_at

  constructor(state?: Partial<State>) {
    if (state) {
      Object.assign(this, state);
    }
  }
}

export function newState(state: State | Record<string, any> | null): State {
  if (state !== null) {
    if (state instanceof State) {
      return cloneDeep(state);
    }
    return new State(state);
  }
  return new State();
}

// 表单验证规则

// 表格搜索表单

// 表格列
export const columns = [
  {
    title: 'id',
    key: 'id',
    align: 'left',
    width: 80,
  },
  {
    title: '名称',
    key: 'name',
    align: 'left',
    width: 120,
  },
  {
    title: '结算方式 ',
    key: 'type',
    align: 'left',
    width: 150,
    render(row) {
      if (isNullObject(row.type)) {
        return ``;
      }
      return h(
        NTag,
        {
          style: {
            marginRight: '6px',
          },
          type: getOptionTag(options.value.settlement_type, row.type),
          bordered: false,
        },
        {
          default: () => getOptionLabel(options.value.settlement_type, row.type),
        }
      );
    },
  },
  {
    title: '结算周期',
    key: 'cycle',
    align: 'left',
    width: 100,
    render(row) {
      if (isNullObject(row.cycle) || row.type == 1) {
        return `--`;
      }
      return h(
        NTag,
        {
          style: {
            marginRight: '6px',
          },
          type: getOptionTag(options.value.settlement_citcle, row.cycle),
          bordered: false,
        },
        {
          default: () => getOptionLabel(options.value.settlement_citcle, row.cycle),
        }
      );
    },
  },
  {
    title: '结算比例',
    key: 'rate',
    align: 'left',
    width: 80,
    render(row){
      if(row.type == 1){
        return `--`;
      }else{
        return row.rate + '%'
      }
    }
  },
  {
    title: '结算成本控制',
    key: 'cost',
    align: 'left',
    width: 150,
    render(row){
      if(row.cost){
        var htmlArr = [];
        row.cost.split(',').forEach((item) => {
          var htmlItem = h(
            'div',
            {
              style: {
                marginTop: '10px'
              }
            },
            [
              h(
                NTag,
                {
                  style: {
                    marginRight: '6px',
                  },
                  type: getOptionTag(options.value.settlement_cost, parseInt(item)),
                  bordered: false,
                },
                {
                  default: () => getOptionLabel(options.value.settlement_cost, parseInt(item)),
                }
              )
            ]
          )
          htmlArr.push(htmlItem)
        })
      }else{
        var htmlArr = `--`;
      }
      if(row.type == 1){
        return `--`;
      }else{
        return h(
          'div',
          null,
          htmlArr
        )
      }
    }
  },
  {
    title: '状态',
    key: 'status',
    align: 'left',
    width: 80,
    render(row) {
      if (isNullObject(row.status)) {
        return ``;
      }
      return h(
        NTag,
        {
          style: {
            marginRight: '6px',
          },
          type: getOptionTag(options.value.sys_normal_disable, row.status),
          bordered: false,
        },
        {
          default: () => getOptionLabel(options.value.sys_normal_disable, row.status),
        }
      );
    },
  },
  // {
  //   title: '提现审核',
  //   key: 'isAuditWithdraw',
  //   align: 'left',
  //   width: 80,
  //   render(row) {
  //     if (isNullObject(row.isAuditWithdraw)) {
  //       return ``;
  //     }
  //     return h(
  //       NTag,
  //       {
  //         style: {
  //           marginRight: '6px',
  //         },
  //         type: getOptionTag(options.value.sys_switch, row.isAuditWithdraw),
  //         bordered: false,
  //       },
  //       {
  //         default: () => getOptionLabel(options.value.sys_switch, row.isAuditWithdraw),
  //       }
  //     );
  //   },
  // },
  {
    title: '创建时间',
    key: 'createAt',
    align: 'left',
    width: 180,
  },
  {
    title: '更新时间',
    key: 'updateAt',
    align: 'left',
    width: 180,
  },
];

// 字典数据选项
export const options = ref({
  sys_normal_disable: [] as Option[],
  sys_switch: [] as Option[],
  settlement_type: [] as Option[],
  settlement_citcle: [] as Option[],
  settlement_cost: [] as Option[],
});

// 加载字典数据选项
export function loadOptions() {
  Dicts({
    types: ['sys_normal_disable', 'sys_switch', 'settlement_type', 'settlement_citcle', 'settlement_cost'],
  }).then((res) => {
    options.value = res;
  });
}


