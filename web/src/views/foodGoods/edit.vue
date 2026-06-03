<template>
  <div>
    <n-drawer v-model:show="showModal" :width="dialogWidth" :z-index="99">
      <n-drawer-content closable :header-style="{
                          padding: '20px',
                        }" :body-content-style="{
                          padding: '20px',
                        }" :footer-style="{
                          padding: '12px 20px',
                        }"> 
            <template #header>
              <div style="font-weight: 500;font-size: 18px;color: #3D3D3D;line-height: 25px;">{{ formValue.id > 0 ? (restaurantName ? restaurantName + '-编辑套餐 #' + formValue.id : '编辑套餐 #' + formValue.id) : (restaurantName ? restaurantName + '-添加套餐' : '添加套餐') }}</div>
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
                ref="formRef"
                :model="formValue"
                :rules="rules"
                label-placement="top"
                label-width="auto"
              >
                <n-tabs type="line" animated v-model:value="tabValue">
                  <n-tab-pane name="base" tab="基础设置">
                    <n-grid :cols="1">
                      <n-gi span="1" v-if="rCooperateTypeId == 4">
                        <n-form-item label="Toreta课程ID" path="toretaCourseId">
                          <n-input placeholder="Toreta课程ID" v-model:value="formValue.toretaCourseId" style="width: 400px"/>
                        </n-form-item>
                      </n-gi>
                      <n-gi span="1">
                        <n-form-item label="套餐名称_简体中文" path="goodsName_zh" :show-require-mark="true">
                          <n-input placeholder="请输入套餐名称" v-model:value="nameLanguage.zh" @blur="translate('name','zh',nameLanguage.zh)" style="width: 400px"/>
                        </n-form-item>
                      </n-gi>
                      <n-gi span="1">
                        <n-form-item label="套餐名称_English" path="goodsName_en">
                          <n-input placeholder="请输入套餐名称" v-model:value="nameLanguage.en" @blur="translate('name','en',nameLanguage.en)" style="width: 400px"/>
                        </n-form-item>
                      </n-gi>
                      <n-gi span="1">
                        <n-form-item label="套餐名称_日本语" path="goodsName_ja">
                          <n-input placeholder="请输入套餐名称" v-model:value="nameLanguage.ja" style="width: 400px"/>
                        </n-form-item>
                      </n-gi>
                      <n-gi span="1">
                        <n-form-item label="促销语" path="introduction">
                          <n-input type="textarea" placeholder="促销语" v-model:value="formValue.introduction" style="width: 400px"/>
                        </n-form-item>
                      </n-gi>
                      <!--            <n-gi span="1">
                                    <n-form-item label="标签" path="labelIdsArr">
                                      <n-checkbox-group v-model:value="formValue.labelIdsArr">
                                        <n-space>
                                          <n-checkbox
                                            v-for="item in labelsArr"
                                            :key="item.id"
                                            :value="item.id"
                                            :label="item.name"
                                          />
                                        </n-space>
                                      </n-checkbox-group>
                                    </n-form-item>
                                  </n-gi>-->
                      <n-gi span="1">
                        <n-form-item label="图集" path="images" :show-require-mark="true" :show-feedback='false'>
                          <FileChooser :maxNumber="10" v-model:value="formValue.imagesArr" />
                        </n-form-item>
                        <div style="color:#9EA4AA; font-size: 12px;line-height: 34px;margin: 0 0 24px;">建议尺寸：686px*324px</div>
                      </n-gi>
                      <n-gi span="1">
                        <n-form-item label="套餐售价" path="price">
                          <n-input-group>
                            <n-input-number placeholder="请输入套餐售价" :min="0" v-model:value="formValue.price" style="width: 100px" />
                            <n-input-group-label>JPY</n-input-group-label>
                          </n-input-group>
                        </n-form-item>
                      </n-gi>
                      <n-gi span="1">
                        <n-form-item label="套餐原价" path="marketPrice">
                          <n-input-group>
                            <n-input-number placeholder="请输入套餐原价" :min="0" v-model:value="formValue.marketPrice" style="width: 100px" />
                            <n-input-group-label>JPY</n-input-group-label>
                          </n-input-group>
                        </n-form-item>
                      </n-gi>
                      <n-gi span="1">
                        <n-form-item label="单次最大可预定数" path="maxOrderNum">
                          <n-input-group>
                            <n-input-number placeholder="请输入" :min="0" v-model:value="formValue.maxOrderNum" style="width: 100px" />
                            <n-input-group-label>单</n-input-group-label>
                          </n-input-group>
                        </n-form-item>
                      </n-gi>
                      <n-gi span="1">
                        <n-form-item label="用餐停留时间" path="timeDuration" :show-require-mark="true">
                          <n-radio-group v-model:value="formValue.timeDuration" name="timeDuration">
                            <n-radio-button
                              v-for="item in timeDurationArr"
                              :key="item.value"
                              :value="item.value"
                              :label="item.label"
                            />
                          </n-radio-group>
                        </n-form-item>
                      </n-gi>
                      <n-gi span="1">
                        <n-form-item label="套餐详情_简体中文" path="goodsContent_zh" :show-require-mark="true">
                          <n-input type="textarea" placeholder="套餐详情" v-model:value="descriptionLanguage.zh" @blur="translate('description','zh',descriptionLanguage.zh)" style="width: 400px"/>
                        </n-form-item>
                      </n-gi>
                      <n-gi span="1">
                        <n-form-item label="套餐详情_English" path="goodsContent_en">
                          <n-input type="textarea" placeholder="套餐详情" v-model:value="descriptionLanguage.en" @blur="translate('description','en',descriptionLanguage.en)" style="width: 400px"/>
                        </n-form-item>
                      </n-gi>
                      <n-gi span="1">
                        <n-form-item label="套餐详情_日本语" path="goodsContent_ja">
                          <n-input type="textarea" placeholder="套餐详情" v-model:value="descriptionLanguage.ja" style="width: 400px"/>
                        </n-form-item>
                      </n-gi>
                      <n-gi span="1">
                        <n-form-item label="注意事项" path="notice">
                          <n-input type="textarea" placeholder="注意事项" v-model:value="formValue.notice" style="width: 400px"/>
                        </n-form-item>
                      </n-gi>
                      <n-gi span="1">
                        <n-form-item label="状态" path="goodsState">
                          <n-radio-group v-model:value="formValue.goodsState" name="goodsState">
                            <n-radio-button
                              :value="1"
                              label="立刻上架"
                            />
                            <n-radio-button
                              :value="2"
                              label="放入仓库"
                            />
                          </n-radio-group>
                        </n-form-item>
                      </n-gi>
                      <n-gi span="1">
                        <n-form-item label="礼品券兑换专属" path="isThCouponExclusive">
                          <n-switch v-model:value="formValue.isThCouponExclusive" :unchecked-value="0" :checked-value="1"/>
                        </n-form-item>
                      </n-gi>
                      <n-gi span="1" v-if="formValue.isThCouponExclusive == 1">
                        <n-form-item label="绑定礼品券" path="thCouponId">
                          <n-select
                            v-model:value="formValue.thCouponId"
                            :options="couponList"
                            label-field="couponName"
                            value-field="id"
                            style="width: 200px"
                          />
                        </n-form-item>
                      </n-gi>
                      <n-gi span="1">
                        <n-form-item label="无需支付" path="isNoPay">
                          <n-switch v-model:value="formValue.isNoPay" :unchecked-value="0" :checked-value="1"/>
                        </n-form-item>
                      </n-gi>
                    </n-grid>
                  </n-tab-pane>
                  <n-tab-pane name="order" tab="预定设置">
                    <n-grid :cols="1">
                      <n-gi span="1">
                        <n-form-item label="时间段限制" path="maxTimeOrderOpen">
                          <n-radio-group v-model:value="formValue.maxTimeOrderOpen" name="maxTimeOrderOpen">
                            <n-space>
                              <n-radio :value="1">
                                开启
                              </n-radio>
                              <n-radio :value="2">
                                关闭
                              </n-radio>
                            </n-space>
                          </n-radio-group>
                        </n-form-item>
                        <n-form-item :label="index == 0 ? '时间段' : ' '" :show-label="index == 0 ? true : false" :path="'orderTime1'+index" v-for="(item, index) in orderTimeForm1Arr" :key="index" v-if="formValue.maxTimeOrderOpen == 1">
                          <n-time-picker placeholder="开始" format="HH:mm" v-model:formatted-value="item.startTime" />
                          <span style="margin: 0 5px">~</span>
                          <n-time-picker placeholder="结束" format="HH:mm" v-model:formatted-value="item.endTime" />
                          <n-input-group style="margin-left: 5px">
                            <n-input-number placeholder="单数" :min="0" :precision="0" :show-button="false" v-model:value="item.num" style="width: 100px" />
                            <n-input-group-label style="margin-right: 12px">单</n-input-group-label>
                            <n-button text type="success" style="margin-left: 12px" attr-type="button" @click="addItem1" v-if="index == 0">
                              <template #icon>
                                <n-icon>
                                  <PlusCircleOutlined />
                                </n-icon>
                              </template>
                            </n-button>
                            <n-button text type="error" style="margin-left: 12px" attr-type="button" @click="removeItem1(index)" v-if="index > 0">
                              <template #icon>
                                <n-icon>
                                  <MinusCircleOutlined />
                                </n-icon>
                              </template>
                            </n-button>
                          </n-input-group>

                        </n-form-item>

                        <!-- 时间段说明提示 -->
                        <div>
                          <span style="color: #e74c3c; font-size: 12px;">
                            注意：时间段包含开始时间，不包含结束时间。例如：00:00~06:00 表示 00:00≤时间&lt;06:00
                          </span>
                        </div>
                      </n-gi>
                    </n-grid>
                  </n-tab-pane>
                </n-tabs>

              </n-form>
            </n-spin>
        </n-drawer-content>
    </n-drawer>
  </div>
</template>

<script lang="ts" setup>
import {ref, onMounted, reactive, computed} from 'vue';
import { Edit, View } from '@/api/foodGoods';
import { State, newState, rules } from './model';
import { useMessage } from 'naive-ui';
import {List as labelList} from "@/api/foodLabel";
import {View as RestaurantView} from "@/api/foodRestaurant";
import {List as CouponList} from "@/api/thCoupon";
import {Text} from "@/api/translate";
import {jsontoobj} from "@/utils/smjcomm";
import {adaModalWidth} from "@/utils/hotgo";
import {MinusCircleOutlined, PlusCircleOutlined} from "@vicons/antd";

const emit = defineEmits(['reloadTable']);
const initRestaurantId = ref<number>(0)
const tabValue = ref('base')
const restaurantName = ref('');
const rCooperateTypeId = ref('');
const message = useMessage();
const loading = ref(false);
const showModal = ref(false);
const formValue = ref<State>(newState(null));
const formRef = ref<any>({});
const formBtnLoading = ref(false);
const dialogWidth = computed(() => {
  return adaModalWidth(650);
});
const orderTimeForm1Arr = ref([{ startTime: null, endTime: null, num: null }]);
const couponList = ref([]);
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
const labelsArr = ref([])
const timeDurationArr = reactive([
  {
    label: '无限制',
    value: 0
  },
  {
    label: '15分',
    value: 15
  },
  {
    label: '30分',
    value: 30
  },
  {
    label: '60分',
    value: 60
  },
  {
    label: '90分',
    value: 90
  },
  {
    label: '120分',
    value: 120
  },
])

function translate(type,lang,text){
  if(type == 'name'){
    if(lang == 'zh'){
      // 简体中文
      nameLanguage.zh_CN = ''
    }else if(lang == 'en'){
      // 韩语
      nameLanguage.ko = ''
    }
  }else if(type == 'description'){
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
      }else if(type == 'description'){
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

const removeItem1 = (index: number) => {
  orderTimeForm1Arr.value.splice(index, 1)
}

const addItem1 = () => {
  orderTimeForm1Arr.value.push({ startTime: null, endTime: null, num: null })
}

async function loadCouponList(){
  let couponArrListOrg = await CouponList({
    Pagination: false
  })
  couponList.value = couponArrListOrg.list
}

function confirmForm(e) {

  if(formValue.value.isThCouponExclusive == 1 && !formValue.value.thCouponId){
    message.error('请绑定礼品券');
    return false;
  }

  e.preventDefault();
  formBtnLoading.value = true;
  formRef.value.validate((errors) => {
    if (!errors) {
      if(formValue.value.imagesArr.length <= 0){
        formBtnLoading.value = false;
        message.error('请上传图集');
        return false;
      }
      formValue.value.labelIds = formValue.value.labelIdsArr.join(',')
      formValue.value.images = formValue.value.imagesArr.join(',')
      formValue.value.nameLanguage = nameLanguage
      formValue.value.descriptionLanguage = descriptionLanguage
      formValue.value.orderTimeForm = JSON.stringify(orderTimeForm1Arr.value)
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

defineExpose({
  openModal,
});

async function loadLabelList(){
  const labelsArrList = await labelList({
    type: 2,
    status: 1,
    Pagination: false
  })
  labelsArr.value = labelsArrList.list
}

async function loadRestaurantView(){
  let info = await RestaurantView({ id: initRestaurantId.value })
  restaurantName.value = info.name
  rCooperateTypeId.value = info.cooperateTypeId
}

async function load(id) {
  const detailRes = await View({
    id: id,
    isLanguage: true
  })
  formValue.value = detailRes
  if(detailRes.isThCouponExclusive == 1){
    formValue.value.thCouponId = detailRes.thCouponId > 0 ? detailRes.thCouponId : null;
  }else{
    formValue.value.thCouponId = null;
  }
  formValue.value.labelIdsArr = detailRes.labelIds ? detailRes.labelIds.split(',').map((item)=>{
    return parseInt(item)
  }) : [];
  formValue.value.imagesArr = detailRes.images ? detailRes.images.split(',').map((item)=>{
    return item
  }) : [];
  formValue.value.timeDuration = detailRes.timeDuration ? parseInt(detailRes.timeDuration) : 0

  if(detailRes.nameLanguage){
    const nameLanguageObj = jsontoobj(detailRes.nameLanguage);
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
  if(detailRes.descriptionLanguage){
    const descriptionLanguageObj = jsontoobj(detailRes.descriptionLanguage);
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

  orderTimeForm1Arr.value = detailRes.orderTimeForm ? JSON.parse(detailRes.orderTimeForm) : [{ startTime: null, endTime: null, num: null }]
}

async function openModal(state: State, restaurantId: number) {
  showModal.value = true;
  loading.value = true;

  tabValue.value = 'base'

  initRestaurantId.value = restaurantId;

  await loadLabelList()
  await loadRestaurantView()
  await loadCouponList()


  // 新增
  if (!state || state.id < 1) {
    formValue.value = newState(state);
    formValue.value.restaurantId = restaurantId;
    nameLanguage.zh = '';
    nameLanguage.en = '';
    nameLanguage.ja = '';
    nameLanguage.ko = '';
    nameLanguage.zh_CN = '';

    descriptionLanguage.zh = '';
    descriptionLanguage.en = '';
    descriptionLanguage.ja = '';
    descriptionLanguage.ko = '';
    descriptionLanguage.zh_CN = '';

    loading.value = false;
    return;
  }

  // 编辑
  await load(state.id);
  loading.value = false;
}
</script>

<style lang="less"></style>


