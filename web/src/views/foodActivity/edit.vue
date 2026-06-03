<template>
  <div>
    <n-spin :show="show" description="请稍候...">
      <n-card
        :bordered="false"
        class="proCard mt-4"
        size="small"
        :segmented="{ content: true }"
        :title="formValue.id > 0 ? '编辑订餐活动 #' + formValue.id : '添加订餐活动'"
      >
        <n-form
          ref="formRef"
          :model="formValue"
          :rules="rules"
          :label-placement="settingStore.isMobile ? 'top' : 'left'"
          :label-width="200"
          class="py-4"
        >
          <n-grid cols="1 s:1 m:1 l:1 xl:1 2xl:1" responsive="screen">
            <n-gi span="1">
              <n-form-item label="活动日期" path="date">
                <DatePicker v-model:formValue="formValue.date" type="date" style="width: 230px"
                            :is-date-disabled="disablePreviousDate"/>
              </n-form-item>
            </n-gi>
            <n-gi span="1">
              <n-form-item label="标题_简体中文" path="name_zh">
                <n-input placeholder="请输入餐厅名称" v-model:value="nameLanguage.zh" @blur="translate('name','zh',nameLanguage.zh)" style="width: 250px" />
              </n-form-item>
            </n-gi>
            <n-gi span="1">
              <n-form-item label="标题_English" path="name_en">
                <n-input placeholder="请输入餐厅名称" v-model:value="nameLanguage.en" @blur="translate('name','en',nameLanguage.en)" style="width: 250px" />
              </n-form-item>
            </n-gi>
            <n-gi span="1">
              <n-form-item label="标题_日本语" path="name_ja">
                <n-input placeholder="请输入餐厅名称" v-model:value="nameLanguage.ja" style="width: 250px" />
              </n-form-item>
            </n-gi>
<!--            <n-gi span="1">
              <n-form-item label="副标题" path="subName">
                <n-input placeholder="请输入副标题" v-model:value="formValue.subName" style="width: 35%"/>
              </n-form-item>
            </n-gi>-->
            <n-gi span="1">
              <n-form-item label="图片" path="pic">
                <!--                  <n-input placeholder="请输入图片" v-model:value="formValue.pic" />-->
                <UploadImage :maxNumber="1" v-model:value="formValue.pic" />
              </n-form-item>
            </n-gi>
            <n-gi span="1">
              <n-form-item label="状态" path="status">
                <!--                  <n-select v-model:value="formValue.status" :options="options.sys_normal_disable" />-->
                <n-radio-group v-model:value="formValue.status" name="status">
                  <n-radio-button
                    v-for="status in options.sys_normal_disable"
                    :key="status.value"
                    :value="status.value"
                    :label="status.label"
                  />
                </n-radio-group>
              </n-form-item>
            </n-gi>
            <n-gi>
              <div style="margin-left: 40px">
                <n-button type="primary" @click="chooseCoupon">选择餐厅</n-button>
              </div>
              <div style="margin-left: 40px;margin-top: 10px;margin-bottom: 20px">
                <n-table>
                  <thead>
                  <tr>
                    <th>餐厅名称</th>
                    <th>营业类型</th>
                    <th>地址信息</th>
                    <th>餐厅状态</th>
                    <th>操作</th>
                  </tr>
                  </thead>
                  <tbody>
                  <tr v-for="(item, index) of couponTypeArr" :key="index">
                    <td>{{ item.name }}</td>
                    <td>
                      {{ item.cooperateTypeDetail.typeName }}
                    </td>
                    <td>{{ item.detailAddress }}</td>
                    <td>
                      <n-tag :type="getOptionTag(RestaurantOptions.open_status, item?.openStatus)" size="small" class="min-left-space">
                        {{ getOptionLabel(RestaurantOptions.open_status, item?.openStatus) }}
                      </n-tag>
                    </td>
                    <td>
                      <n-button type="error" @click="handleDeleteCoupon(index)">删除</n-button>
                    </td>
                  </tr>
                  </tbody>
                </n-table>
              </div>
            </n-gi>
          </n-grid>
          <div style="text-align: center">
            <n-space justify="center">
              <n-button type="info" :loading="formBtnLoading" @click="confirmForm">
                确定
              </n-button>
            </n-space>
          </div>
        </n-form>
      </n-card>
    </n-spin>
    <ChooseCoupon ref="chooseCouponRef" @reloadCouponList="chooseCouponInfo"/>
  </div>
</template>

<script lang="ts" setup>
import {ref, onMounted} from 'vue';
import { Edit, View } from '@/api/foodActivity';
import {State, newState, options, loadOptions} from './model';
import DatePicker from '@/components/DatePicker/datePicker.vue';
import { useProjectSettingStore } from '@/store/modules/projectSetting';
import {FormItemRule, useMessage} from 'naive-ui';
import {useRouter} from "vue-router";
import {useTabsViewStore} from "@/store/modules/tabsView";
import {All} from "@/api/foodRestaurant";
import {jsontoobj} from "@/utils/smjcomm";
import UploadImage from "@/components/Upload/uploadImage.vue";
import ChooseCoupon from "@/components/ChooseFoodRestaurant/chooseFoodRestaurant.vue";
import {Text} from "@/api/translate";
import {getOptionLabel, getOptionTag} from "@/utils/hotgo";
import {loadOptions as RestaurantLoadOptions, options as RestaurantOptions} from "@/views/foodRestaurant/model";
import {subDays} from "date-fns/esm";

const show = ref(false);
const router = useRouter();
const params = router.currentRoute.value.params;
const formValue = ref<State>(newState(null));

const nameLanguage = ref({
  en: '',
  zh: '',
  ja: '',
  ko: '',
  zh_CN: '',
});


// 表单验证规则
const rules = {

};
const message = useMessage();
const settingStore = useProjectSettingStore();
const formRef = ref<any>({});
const formBtnLoading = ref(false);
const tabsViewStore = useTabsViewStore();
const chooseCouponRef = ref();
const couponTypeIdsArr = ref([]);
const couponTypeArr = ref([])

function  disablePreviousDate(ts: number) {
  // console.log('ts',ts)
  // console.log('subDays(Date.now(), 1)',subDays(Date.now(), 1))
  // console.log('subDays(Date.now(), 1).valueOf()',subDays(Date.now(), 1).valueOf())
  return ts < subDays(Date.now(), 1).valueOf();
}

// 翻译
function translate(type,lang,text){
  if(type == 'name'){
    if(lang == 'zh'){
      // 简体中文
      nameLanguage.value.zh_CN = ''
    }else if(lang == 'en'){
      // 韩语
      nameLanguage.value.ko = ''
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
          nameLanguage.value.zh_CN = _res.zh_CN
        }else if(lang == 'en'){
          // 韩语
          nameLanguage.value.ko = _res.ko
        }
      }
    }).catch((err) => {

    });
  }
}

function chooseCoupon(){
  chooseCouponRef.value.openModal(couponTypeIdsArr.value);
  // chooseCouponRef.value.openModal([]);
}
function handleDeleteCoupon(index){
  couponTypeArr.value.splice(index, 1);
  couponTypeIdsArr.value.splice(index, 1);
  formValue.value.restaurantIds = couponTypeIdsArr.value.join(',')
}

function chooseCouponInfo(couponInfo){
  couponTypeArr.value.push(couponInfo)
  couponTypeIdsArr.value.push(couponInfo.id)
  formValue.value.restaurantIds = couponTypeIdsArr.value.join(',')
}

function confirmForm(e) {
  // 订餐活动名称-多语言
  formValue.value.nameLanguage = nameLanguage.value

  e.preventDefault();

  formBtnLoading.value = true;
  formRef.value.validate((errors) => {
    if (!errors) {
      Edit(formValue.value).then((_res) => {
        message.success('操作成功');
        formBtnLoading.value = false;
        setTimeout(() => {
          tabsViewStore.closeSignal('2');
          router.push({ name: 'pmsActivityIndex', params: {  } });
        });
      }).catch((err) => {
        formBtnLoading.value = false;
      });
    } else {
      message.error('请填写完整信息');
      formBtnLoading.value = false;
    }
  });
}

function getRestaurantList(){
  All({
    restaurantIds: formValue.value.restaurantIds
  }).then((res) => {
    couponTypeArr.value = res.list;
  })
    .finally(() => {
      show.value = false;
    });
}


onMounted(() => {
  show.value = true;
  loadOptions()
  RestaurantLoadOptions()
  if(params.id){
    View({ id: params.id })
      .then((res) => {
        formValue.value = res;
        // couponTypeIdsArr.value = res.restaurantIds ? res.restaurantIds.split(",").map(item => Number(item)) : [];


        if(res.nameLanguage){
          res.nameLanguage = jsontoobj(res.nameLanguage);
          nameLanguage.value.zh = res.nameLanguage.zh && res.nameLanguage.zh.content ? res.nameLanguage.zh.content : ''
          nameLanguage.value.en = res.nameLanguage.en && res.nameLanguage.en.content ? res.nameLanguage.en.content : ''
          nameLanguage.value.ko = res.nameLanguage.ko && res.nameLanguage.ko.content ? res.nameLanguage.ko.content : ''
          nameLanguage.value.ja = res.nameLanguage.ja && res.nameLanguage.ja.content ? res.nameLanguage.ja.content : ''
          if(res.nameLanguage.zh_CN){
            nameLanguage.value.zh_CN = res.nameLanguage.zh_CN.content ? res.nameLanguage.zh_CN.content : ''
          }else if(res.nameLanguage.zh_cn){
            nameLanguage.value.zh_CN = res.nameLanguage.zh_cn.content ? res.nameLanguage.zh_cn.content : ''
          }else{
            nameLanguage.value.zh_CN = ''
          }
        }

        // couponTypeArr.value = res.RestaurantList.map(item => item.RestaurantDetail);

        res.restaurantList.forEach((item) => {
          couponTypeIdsArr.value.push(item.restaurantId)
        })

        formValue.value.restaurantIds = couponTypeIdsArr.value.join(',')
        getRestaurantList()

      })
      .finally(() => {
        show.value = false;
      });
    return;
  }
  show.value = false;
});
</script>

<style lang="less"></style>


