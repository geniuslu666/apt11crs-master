import { ref } from 'vue';
import { cloneDeep } from 'lodash-es';
import { FormSchema } from '@/components/Form';
import {Option} from "@/utils/hotgo";
import { Dicts } from '@/api/dict/dict';

export class State {
  public id = 0;
  public title = '';
  public titleLanguage = { en: '', zh: '', ja: '', ko: '', zh_CN: '' };
  public subTitle = '';
  public subTitleLanguage = { en: '', zh: '', ja: '', ko: '', zh_CN: '' };
  public contentZh = '';
  public contentEn = '';
  public contentJa = '';
  public contentKo = '';
  public contentTw = '';
  public listImage = '';
  public carouselImages = [];
  public dailyCapacity = 10;
  public stock = 0;
  public price = 0;
  public meetingPlace = '';
  public meetingTime = null;
  public maxBookDays = 30;
  public advanceBookDays = 1;
  public tripPlanningZh = '';
  public tripPlanningEn = '';
  public tripPlanningJa = '';
  public tripPlanningKo = '';
  public tripPlanningTw = '';
  public bookingNotesZh = '';
  public bookingNotesEn = '';
  public bookingNotesJa = '';
  public bookingNotesKo = '';
  public bookingNotesTw = '';
  public status = 1;
  public sort = 0;
  public ggLat = '';
  public ggLng = '';
  public contactMobile = '';
  public skuList = {

  };

  constructor(state?: Partial<State>) {
    if (state) {
      Object.assign(this, state);
    }
  }
}

export function newState(state: State | Record<string, any> | null): State {
  if (state !== null) {
    if (state instanceof State) return cloneDeep(state);
    return new State(state);
  }
  return new State();
}

export const rules = {
  listImage: { required: true, trigger: ['blur', 'input'], message: '请上传列表图' },
  price: { required: true, trigger: ['blur', 'input'], type: 'number', message: '请输入售价' },
  meetingPlace: { required: true, trigger: ['blur', 'input'], message: '请输入集合地点' },
  meetingTime: { required: true, trigger: ['blur', 'input'], message: '请选择集合时间' },
  dailyCapacity: { required: true, trigger: ['blur', 'input'], type: 'number', message: '请输入每日接待人数' },
};

export const schemas = ref<FormSchema[]>([
  {
    field: 'title',
    component: 'NInput',
    label: '产品标题',
    componentProps: { placeholder: '请输入产品标题' },
  },
  // {
  //   field: 'status',
  //   component: 'NSelect',
  //   label: '状态',
  //   componentProps: {
  //     placeholder: '请选择状态',
  //     options: [],
  //     onUpdateValue: (e: any) => {
  //       console.log(e);
  //     },
  //   },
  // },
]);

export const options = ref({
  sys_normal_disable: [] as Option[],
});

// 加载字典数据选项
export function loadOptions() {
  Dicts({
    types: ['sys_normal_disable', 'driver_work_status'],
  }).then((res) => {
    options.value = res;
    // for (const item of schemas.value) {
    //   switch (item.field) {
    //     case 'status':
    //       item.componentProps.options = options.value.sys_normal_disable;
    //       break;
    //   }
    // }
  });
}
