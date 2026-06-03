import { cloneDeep } from 'lodash-es';
import {h} from "vue";

export class State {
  public id = 0; // id
  public scene = ''; // 场景
  public condition = null;
  public status = 0; // 状态
  public path = ''; // 路径
  public createAt = ''; // 创建时间
  public updateAt = ''; // 更新时间

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
export const rules = {
};

// 表格列
export const columns = [
  {
    title: 'ID',
    key: 'id',
    align: 'left',
    width: 80,
  },
  {
    title: '搜索条件',
    key: 'condition',
    align: 'left',
    width: 150,
  },
  {
    title: '状态',
    key: 'status',
    align: 'left',
    width: 100,
    render(record){
      if(record.status == 0){
        return h(
          'div',
          {
            style: {
              color: '#919399',
              width: '52px',
              height: '22px',
              background: '#F4F4F5',
              lineHeight: '22px',
              textAlign: 'center',
              fontSize: '14px',
              borderRadius: '2px',
              fontWeight: '400',
            },
            bordered: false,
          },
          {
            default: () => '导出中',
          }
        );
      }else if(record.status == 1){
        return h(
          'div',
          {
            style: {
              color: '#26A763',
              padding: '0 5px',
              height: '22px',
              background: '#E3F4EB',
              lineHeight: '22px',
              textAlign: 'center',
              fontSize: '14px',
              borderRadius: '2px',
              fontWeight: '400',
            },
            bordered: false,
          },
          {
            default: () => '导出成功',
          }
        );
      }else if(record.status == 2){
        return h(
          'div',
          {
            style: {
              color: '#F56C6C',
              padding: '0 5px',
              height: '22px',
              background: '#FEF0F0',
              lineHeight: '22px',
              textAlign: 'center',
              fontSize: '14px',
              borderRadius: '2px',
              fontWeight: '400',
            },
            bordered: false,
          },
          {
            default: () => '导出失败',
          }
        );
      }
    }
  },
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



