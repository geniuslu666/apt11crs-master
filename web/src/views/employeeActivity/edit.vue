<template>
  <div>
    <n-drawer v-model:show="showModal" :width="dialogWidth" @after-leave="afterLeaveCallback">
      <n-drawer-content closable :header-style="{
                    padding: '20px',
                  }" :body-content-style="{
                    padding: '20px',
                  }" :footer-style="{
                    padding: '12px 20px',
                  }">
        <template #header>
          <div style="font-weight: 500;font-size: 18px;color: #3D3D3D;line-height: 25px;">{{ formValue.id > 0 ? '编辑活动 #' + formValue.id : '添加活动' }}</div>
        </template>
        <template #footer>
          <n-button @click="closeForm" style="width: 70px;height: 35px;margin-right: 10px">
            取消
          </n-button>
          <n-button type="info" :loading="formBtnLoading" @click="confirmForm" style="width: 70px;height: 35px;">
            保存
          </n-button>
        </template>
        <n-spin :show="loading" description="请稍候...">
          <n-form
            :model="formValue"
            :rules="rules"
            ref="formRef"
            label-placement="top"
            label-width="auto"
          >
            <n-form-item label="活动名称_简体中文" path="name_zh" :show-require-mark="true">
              <n-input placeholder="请输入活动名称" v-model:value="nameLanguage.zh" style="width: 300px" @blur="translate('name','zh',nameLanguage.zh)"/>
            </n-form-item>
            <n-form-item label="活动名称_English" path="name_en" :show-require-mark="true">
              <n-input placeholder="请输入活动名称" v-model:value="nameLanguage.en" style="width: 300px" @blur="translate('name','en',nameLanguage.en)"/>
            </n-form-item>
            <n-form-item label="活动名称_日本语" path="name_ja" :show-require-mark="true">
              <n-input placeholder="请输入活动名称" v-model:value="nameLanguage.ja" style="width: 300px" />
            </n-form-item>
            <n-form-item label="活动封面" path="cover" :show-feedback='false'>
              <UploadImage :maxNumber="1" v-model:value="formValue.cover" />
            </n-form-item>
            <div style="color:#9EA4AA; font-size: 12px;line-height: 34px;margin: 0 0 24px;">建议尺寸：180x*230px</div>
            <n-form-item label="有效期类型" path="validityType" v-if='false'>
              <n-radio-group v-model:value="formValue.validityType">
                <n-space>
                  <n-radio :value="1">
                    固定时间段
                  </n-radio>
                  <n-radio :value="2">
                    长期有效
                  </n-radio>
                </n-space>
              </n-radio-group>
            </n-form-item>
            <n-form-item label="活动开始时间" path="startTime" :show-require-mark="true" v-if="formValue.validityType == 1">
              <DatePicker v-model:formValue="formValue.startTime" type="date" style="width: 300px"
                                :is-date-disabled="disablePreviousDate" :disabled="formValue.id > 0"/>
            </n-form-item>
            <n-form-item label="活动结束时间" path="endTime" :show-require-mark="true" v-if="formValue.validityType == 1">
              <DatePicker v-model:formValue="formValue.endTime" type="date" style="width: 300px"
                                :is-date-disabled="disablePreviousDate" :disabled="formValue.id > 0"/>
            </n-form-item>
            <n-form-item label="活动描述_简体中文" path="description_zh">
              <n-input type="textarea" placeholder="请输入简介" v-model:value="descriptionLanguage.zh" @blur="translate('description','zh',descriptionLanguage.zh)" style="width: 500px" />
            </n-form-item>
            <n-form-item label="活动描述_English" path="description_en">
              <n-input type="textarea" placeholder="请输入简介" v-model:value="descriptionLanguage.en" @blur="translate('description','en',descriptionLanguage.en)" style="width: 500px" />
            </n-form-item>
            <n-form-item label="活动描述_日本语" path="description_ja">
              <n-input type="textarea" placeholder="请输入简介" v-model:value="descriptionLanguage.ja" style="width: 500px" />
            </n-form-item>
            <n-form-item label="活动规则" path="rule">
              <n-input type="textarea" placeholder="请输入简介" v-model:value="formValue.rule" style="width: 500px" />
            </n-form-item>
            <n-form-item label="礼品券" path="mchIds" :show-require-mark="true">
              <n-button type="primary" @click="chooseThCoupon()">添加礼品券</n-button>
            </n-form-item>
            <n-table style="margin-bottom: 25px">
              <thead>
              <tr>
                <th>礼品券名称</th>
                <th>适用商户</th>
                <th>每人每天领取数(0不限制)</th>
                <th>每人领取数量(>0)</th>
                <th>每人每天核销数(0不限制)</th>
                <th>操作</th>
              </tr>
              </thead>
              <tbody>
              <tr v-for="(item, index) of thCouponArr" :key="index">
                <td>{{ item.couponName }}</td>
                <td>
                  <n-ellipsis expand-trigger="click" :line-clamp="1" :tooltip="false" style="width: 200px">
                    <div v-for="(item1, index1) of item.mchList" :key="index1">{{ item1.mchInfo.name }}</div>
                  </n-ellipsis>
                </td>
                <td>
                  <n-input-number v-model:value="item.perNum" min="0" placeholder="每人每天可领取数" :show-button="false" style="width: 80px"/>
                </td>
                <td>
                    <n-input-number v-model:value="item.num" min="1" placeholder="每人领取数量" :show-button="false" style="width: 80px"/>
                </td>
                <td>
                  <n-input-number v-model:value="item.perVerifyNum" min="0" placeholder="每人每天可核销数" :show-button="false" style="width: 80px"/>
                </td>
                <td>
                  <n-button type="error" @click="handleDeleteThCoupon(item.id)">删除</n-button>
                </td>
              </tr>
              </tbody>
            </n-table>
            <!-- 领取限制，限制周六、周日不可领取、分开勾选 -->
            <n-form-item label="礼品券领取限制" path="saturdayLimit" :show-require-mark="false">
              <n-space>
                <n-checkbox v-model:checked="saturdayLimitChecked">
                  周六不允许领取
                </n-checkbox>
                <n-checkbox v-model:checked="sundayLimitChecked">
                  周日不允许领取
                </n-checkbox>
              </n-space>
            </n-form-item>
            <n-form-item label="券有效期" path="couponValidity">
              <n-switch :unchecked-value="2" :checked-value="1" v-model:value="formValue.couponValidity"/><span style="margin-left: 10px">随活动时间</span>
              <template #feedback> <span style="color: red">不开启则使用券自身有效期</span> </template>
            </n-form-item>
            <n-form-item label="限制类型" path="restrictionType" style="margin-top: 24px">
              <n-radio-group v-model:value="formValue.restrictionType">
                <n-space>
                  <n-radio :value="1">
                    不限制
                  </n-radio>
                  <n-radio :value="2">
                    限制指定部门
                  </n-radio>
                  <n-radio :value="3">
                    限制指定员工
                  </n-radio>
                </n-space>
              </n-radio-group>
            </n-form-item>
            <n-form-item label="部门" path="departmentIds" :show-require-mark="true" v-if="formValue.restrictionType == 2">
              <n-tree-select
                multiple
                cascade
                checkable
                key-field="id"
                label-field="name"
                children-field="children"
                placeholder="请选择部门"
                :options="departmentTreeOptions"
                v-model:value="formValue.departmentIds"
              />
            </n-form-item>
            <template v-if="formValue.restrictionType == 3">
              <n-form-item label="员工" path="employeeIds" :show-require-mark="true">
                <n-button type="primary" @click="chooseEmployee()">选择员工</n-button>
              </n-form-item>
              <n-table style="margin-bottom: 25px">
                <thead>
                <tr>
                  <th>名称</th>
                  <th>电话</th>
                  <th>所属部门</th>
                  <th>操作</th>
                </tr>
                </thead>
                <tbody>
                <tr v-for="(item, index) of employeeArr" :key="index">
                  <td>{{ item.name }}</td>
                  <td>{{ item.phoneArea + ' ' + item.phone }}</td>
                  <td>{{ item.departmentDetail.name }}</td>
                  <td>
                    <n-button type="error" @click="handleDeleteEmployee(item.id)">删除</n-button>
                  </td>
                </tr>
                </tbody>
              </n-table>
            </template>
            <n-form-item label="状态" path="isEnabled">
              <n-radio-group v-model:value="formValue.isEnabled" name="isEnabled">
                <n-radio-button
                  v-for="status in options.sys_normal_disable"
                  :key="status.value"
                  :value="status.value"
                  :label="status.label"
                />
              </n-radio-group>
            </n-form-item>
          </n-form>
        </n-spin>
      </n-drawer-content>
    </n-drawer>

    <ThChooseCoupon ref="chooseThCouponRef" @reloadThCoupon="chooseThCouponInfo"/>
    <ChooseEmployee ref="chooseEmployeeRef" @reloadEmployee="chooseEmployeeInfo"/>
  </div>
</template>

<script lang="ts" setup>
import {computed, ref, reactive} from 'vue';
import { Edit, View } from '@/api/employeeActivity';
import { options, State, newState, rules } from './model';
import { useMessage } from 'naive-ui';
import {adaModalWidth} from "@/utils/hotgo";
import { GetDepartmentTree } from '@/api/employeeDepartment';
import {subDays} from "date-fns/esm";
import { EmployeeInter } from '/#/employee';
import { ThCouponMchInter } from '/#/thCoupon';
import ThChooseCoupon from "./chooseThCoupon.vue";
import ChooseEmployee from "./chooseEmployee.vue";
import {List as EmployeeList} from "@/api/employee";
import {All} from "@/api/thCoupon";
import {Text} from "@/api/translate";
import {jsontoobj} from "@/utils/smjcomm";

const emit = defineEmits(['reloadTable']);
const message = useMessage();
const loading = ref(false);
const showModal = ref(false);
const formValue = ref<State>(newState(null));
const formRef = ref<any>({});
const formBtnLoading = ref(false);
const dialogWidth = computed(() => {
  return adaModalWidth(1100);
});
const departmentTreeOptions = ref([]);

interface SelectThCouponInter {
  id: number;
  couponName: string;
  mchList: ThCouponMchInter[];
  num: number;
  perNum: number;
  perVerifyNum: number;
}

interface ThCouponIdsNumInter {
  couponId: number;
  num: number;
  perNum: number;
  perVerifyNum: number;
}

const thCouponIdsNumArr = ref<ThCouponIdsNumInter[]>([]);
const thCouponArr = ref<SelectThCouponInter[]>([])
const chooseThCouponRef = ref();

const employeeIdsArr = ref<number[]>([]);
const employeeArr = ref<EmployeeInter[]>([])
const chooseEmployeeRef = ref();

const nameLanguage = reactive({
  en: '',
  zh: '',
  ja: '',
  ko: '',
  zh_CN: '',
});

const descriptionLanguage = reactive({
  en: '',
  zh: '',
  ja: '',
  ko: '',
  zh_CN: '',
});

// 添加计算属性确保复选框状态正确
const saturdayLimitChecked = computed({
  get: () => {
    const limits = formValue.value.limitWeek ? formValue.value.limitWeek.split(',') : [];
    return limits.includes('6');
  },
  set: (value) => {
    const limits = formValue.value.limitWeek ? formValue.value.limitWeek.split(',') : [];
    if (value) {
      // 不允许周六 - 添加6
      if (!limits.includes('6')) {
        limits.push('6');
      }
      formValue.value.limitWeek = limits.join(',');
    } else {
      // 允许周六 - 移除6
      formValue.value.limitWeek = limits.filter(limit => limit !== '6').join(',');
    }
  }
});

const sundayLimitChecked = computed({
  get: () => {
    const limits = formValue.value.limitWeek ? formValue.value.limitWeek.split(',') : [];
    return limits.includes('7');
  },
  set: (value) => {
    const limits = formValue.value.limitWeek ? formValue.value.limitWeek.split(',') : [];
    if (value) {
      // 不允许周日 - 添加7
      if (!limits.includes('7')) {
        limits.push('7');
      }
      formValue.value.limitWeek = limits.join(',');
    } else {
      // 允许周日 - 移除7
      formValue.value.limitWeek = limits.filter(limit => limit !== '7').join(',');
    }
  }
});

function translate(type,lang,text){
  if(type == 'name'){
    if(lang == 'zh'){
      // 简体中文
      nameLanguage.zh_CN = ''
    }else if(lang == 'en'){
      // 韩语
      nameLanguage.ko = ''
    }
  }
  if(type == 'description'){
    if(lang == 'zh'){
      // 简体中文
      descriptionLanguage.zh_CN = ''
    }else if(lang == 'en'){
      // 韩语
      descriptionLanguage.ko = ''
    }
  }
  if(text){
    if(lang == 'en'){
      text = text.toLowerCase()
    }
    Text({
      text: text
    }).then((_res) => {
      if(type == 'name'){
        if(lang == 'zh'){
          // 简体中文
          nameLanguage.zh_CN = _res.zh_CN
        }else if(lang == 'en'){
          // 韩语
          nameLanguage.ko = _res.ko
        }
      }
      if(type == 'description'){
        if(lang == 'zh'){
          // 简体中文
          descriptionLanguage.zh_CN = _res.zh_CN
        }else if(lang == 'en'){
          // 韩语
          descriptionLanguage.ko = _res.ko
        }
      }
    }).catch((err) => {

    });
  }
}

function  disablePreviousDate(ts: number) {
  return ts < subDays(Date.now(), 1).valueOf();
}

async function loadDepartmentTree() {
  try {
    const res = await GetDepartmentTree();
    departmentTreeOptions.value = res.list || [];
  } catch (error) {
    console.error('加载部门树失败:', error);
  }
}

async function getInfo(id){
  const res = await View({ id: id, isLanguage: true });
  formValue.value = res;
  if(res.nameLanguage){
    const nameLanguageObj = jsontoobj(res.nameLanguage) as any;
    nameLanguage.zh = nameLanguageObj.zh && nameLanguageObj.zh.content ? nameLanguageObj.zh.content : ''
    nameLanguage.en = nameLanguageObj.en && nameLanguageObj.en.content ? nameLanguageObj.en.content : ''
    nameLanguage.ko = nameLanguageObj.ko && nameLanguageObj.ko.content ? nameLanguageObj.ko.content : ''
    nameLanguage.ja = nameLanguageObj.ja && nameLanguageObj.ja.content ? nameLanguageObj.ja.content : ''
    if(nameLanguageObj.zh_CN){
      nameLanguage.zh_CN = nameLanguageObj.zh_CN.content ? nameLanguageObj.zh_CN.content : ''
    }else if(nameLanguageObj.zh_cn){
      nameLanguage.zh_CN = nameLanguageObj.zh_cn.content ? nameLanguageObj.zh_cn.content : ''
    }else{
      nameLanguage.zh_CN = ''
    }
  }
  if(res.descriptionLanguage){
    const descriptionLanguageObj = jsontoobj(res.descriptionLanguage) as any;
    descriptionLanguage.zh = descriptionLanguageObj.zh && descriptionLanguageObj.zh.content ? descriptionLanguageObj.zh.content : ''
    descriptionLanguage.en = descriptionLanguageObj.en && descriptionLanguageObj.en.content ? descriptionLanguageObj.en.content : ''
    descriptionLanguage.ko = descriptionLanguageObj.ko && descriptionLanguageObj.ko.content ? descriptionLanguageObj.ko.content : ''
    descriptionLanguage.ja = descriptionLanguageObj.ja && descriptionLanguageObj.ja.content ? descriptionLanguageObj.ja.content : ''
    if(descriptionLanguageObj.zh_CN){
      descriptionLanguage.zh_CN = descriptionLanguageObj.zh_CN.content ? descriptionLanguageObj.zh_CN.content : ''
    }else if(descriptionLanguageObj.zh_cn){
      descriptionLanguage.zh_CN = descriptionLanguageObj.zh_cn.content ? descriptionLanguageObj.zh_cn.content : ''
    }else{
      descriptionLanguage.zh_CN = ''
    }
  }
  thCouponIdsNumArr.value = []
  thCouponArr.value = []
  if(res.couponList){
    res.couponList.forEach((item) => {
      thCouponIdsNumArr.value.push({couponId: item.couponId, num: item.availableQuantity, perNum: item.perDayAvailable, perVerifyNum: item.perDayVerify})
    })
  }
  await getThCouponList()
  if(thCouponArr.value.length > 0){
    thCouponArr.value.forEach(m => {
      m.num = thCouponIdsNumArr.value.find(n => n.couponId == m.id)?.num || 1
      m.perNum = thCouponIdsNumArr.value.find(n => n.couponId == m.id)?.perNum || 0
      m.perVerifyNum = thCouponIdsNumArr.value.find(n => n.couponId == m.id)?.perVerifyNum || 0
    })
  }

  if(res.restrictionType == 2){
    formValue.value.departmentIds = res.employeeDepartmentList.map(m => m.departmentId);
  }

  if(res.restrictionType == 3){
    employeeIdsArr.value = res.employeeList.map(m => m.employeeId);
    employeeArr.value = [];
    await getEmployeeList()
  }
}

// 选择礼品券相关
function chooseThCoupon(){
  chooseThCouponRef.value.openModal(thCouponIdsNumArr.value.map(m => m.couponId));
}

function handleDeleteThCoupon(id){
  thCouponArr.value = thCouponArr.value.filter(m => m.id !== id)
  thCouponIdsNumArr.value = thCouponIdsNumArr.value.filter(m => m.couponId !== id)
}

async function chooseThCouponInfo(selectedIds){
  thCouponIdsNumArr.value = selectedIds.map(m => ({couponId: m, num: 1, perNum: 0, perVerifyNum: 1}))
    loading.value = true;
    const oldMap = new Map(thCouponArr.value.map(m => [m.id,m]))
    console.log('oldMap',oldMap)
    await getThCouponList()

    thCouponArr.value = thCouponArr.value.map(m => ({
      ...m,
      num: oldMap.get(m.id)?.num || 1,
      perNum: oldMap.get(m.id)?.perNum || 0,
      perVerifyNum: oldMap.get(m.id)?.perVerifyNum || 1
    }))
    console.log('thCouponArr',thCouponArr.value)
    loading.value = false;
}


async function getThCouponList(){
  const res = await All({
    couponIds: thCouponIdsNumArr.value.map(m => m.couponId).join(',')
  });
  thCouponArr.value = res.list
  console.log('thCouponArr',thCouponArr.value)
}

// 选择员工相关
function chooseEmployee(){
  chooseEmployeeRef.value.openModal(employeeIdsArr.value);
}

function handleDeleteEmployee(id){
  employeeArr.value = employeeArr.value.filter(m => m.id !== id)
  employeeIdsArr.value = employeeIdsArr.value.filter(m => m !== id)
}

async function chooseEmployeeInfo(selectedIds){
  employeeIdsArr.value = selectedIds
    loading.value = true;
    const oldMap = new Map(employeeArr.value.map(m => [m.id,m]))
    console.log('oldMap',oldMap)
    await getEmployeeList()

    employeeArr.value = employeeArr.value.map(m => ({
      ...m
    }))
    console.log('employeeArr',employeeArr.value)
    loading.value = false;
}


async function getEmployeeList(){
  const res = await EmployeeList({
    employeeIds: employeeIdsArr.value.join(','),
    pagination: false
  });
  employeeArr.value = res.list
  console.log('employeeArr',employeeArr.value)
}

async function openModal(state: State) {
  showModal.value = true;
  loading.value = true;

  await loadDepartmentTree();

  // 新增
  if (!state || state.id < 1) {
    formValue.value = newState(state);
    Object.assign(nameLanguage, {
      en: '',
      zh: '',
      ja: '',
      ko: '',
      zh_CN: '',
    })
    Object.assign(descriptionLanguage, {
      en: '',
      zh: '',
      ja: '',
      ko: '',
      zh_CN: '',
    })
    thCouponIdsNumArr.value = []
    thCouponArr.value = []
    employeeIdsArr.value = []
    employeeArr.value = []
    console.log('formValue',formValue.value)
    loading.value = false;
    return;
  }

  await getInfo(state.id)

  loading.value = false;
}

function confirmForm(e) {
  if(!nameLanguage.zh || !nameLanguage.en || !nameLanguage.ja){
    formBtnLoading.value = false;
    message.error('请填写活动名称');
    return false;
  }

  if(!descriptionLanguage.zh || !descriptionLanguage.en || !descriptionLanguage.ja){
    formBtnLoading.value = false;
    message.error('请填写活动描述');
    return false;
  }

  if(formValue.value.validityType == 1 && (!formValue.value.startTime || !formValue.value.endTime)){
    formBtnLoading.value = false;
    message.error('请选择活动时间');
    return false;
  }

  if(thCouponArr.value.length <= 0){
    formBtnLoading.value = false;
    message.error('请选择礼品券');
    return false;
  }

  if(formValue.value.restrictionType == 3 && employeeArr.value.length <= 0){
    formBtnLoading.value = false;
    message.error('请选择员工');
    return false;
  }

  for (let i = 0; i < thCouponArr.value.length; i++) {
    if (thCouponArr.value[i].num <= 0) {
      formBtnLoading.value = false;
      message.error(thCouponArr.value[i].couponName +',请填入写每人可领用数量');
      return false;
    }
  }

  formValue.value.thCouponsArr = thCouponArr.value.map(m => ({
    couponId: m.id,
    availableQuantity: m.num,
    perDayAvailable: m.perNum,
    perDayVerify: m.perVerifyNum,
  }))

  formValue.value.employeeIds = employeeIdsArr.value

  e.preventDefault();
  formBtnLoading.value = true;
  formRef.value.validate((errors) => {
    if (!errors) {
      formValue.value.nameLanguage = nameLanguage
      formValue.value.descriptionLanguage = descriptionLanguage
      Edit(formValue.value).then((_res) => {
        message.success('操作成功');
        formBtnLoading.value = false;
        setTimeout(() => {
          closeForm();
          emit('reloadTable');
        });
      }).catch((err) => {
        formBtnLoading.value = false;
      });
    } else {
      formBtnLoading.value = false;
      message.error('请填写完整信息');
    }
  });
}

function closeForm() {
  showModal.value = false;
  loading.value = false;
}

function afterLeaveCallback() {
  formValue.value = newState(null);
}

defineExpose({
  openModal,
});
</script>

<style lang="less"></style>


