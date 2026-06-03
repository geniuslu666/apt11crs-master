<template>
 <div>
 <n-drawer v-model:show="showModal" :width="dialogWidth">
 <n-drawer-content
 closable
 :header-style="{ padding: '20px' }"
 :body-content-style="{ padding: '20px' }"
 :footer-style="{ padding: '12px 20px' }"
 >
 <template #header>
 <div style="font-weight:500; font-size:18px; color: #3d3d3d; line-height:25px">
 {{ formValue.id >0 ? `编辑核销人员 #${formValue.id}` : '新增核销人员' }}
 </div>
 </template>

 <template #footer>
 <n-button @click="closeForm" style="width:70px; height:35px; margin-right:10px">取消</n-button>
 <n-button type="info" :loading="formBtnLoading" @click="confirmForm" style="width:70px; height:35px">
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
 require-mark-placement="right-hanging"
 >
 <div class="level-detail-div">
 <div class="level-detail-div-title">
 <div></div>
 基本信息
 </div>
 </div>

 <n-grid cols="1600:3" x-gap="80">
 <n-gi>
 <n-form-item label="姓名" path="name">
 <n-input v-model:value="formValue.name" placeholder="请输入姓名" />
 </n-form-item>
 </n-gi>
 <n-gi>
 <n-form-item label="电话" path="mobile">
 <n-input v-model:value="formValue.mobile" placeholder="请输入电话" />
 </n-form-item>
 </n-gi>
 <n-gi>
 <n-form-item label="登录账号" path="username">
 <n-input v-model:value="formValue.username" placeholder="请输入登录账号" />
 </n-form-item>
 </n-gi>
 <n-gi>
 <n-form-item label="密码">
 <n-input
 v-model:value="formValue.password"
 type="password"
 show-password-on="click"
 :placeholder="formValue.id >0 ? '不填则不修改密码' : '请输入密码'"
 />
 </n-form-item>
 </n-gi>
 <n-gi>
 <n-form-item label="状态" path="status">
 <n-radio-group v-model:value="formValue.status" name="status">
 <n-radio-button :value="1" label="启用" />
 <n-radio-button :value="2" label="禁用" />
 </n-radio-group>
 </n-form-item>
 </n-gi>
 </n-grid>

 <div class="level-detail-div" style="margin-top:24px">
 <div class="level-detail-div-title">
 <div></div>
 核销权限范围
 </div>
 </div>

 <n-form-item label="权限模式">
 <n-radio-group v-model:value="formValue.isAllScope">
 <n-space>
 <n-radio :value="true">全部可核销</n-radio>
 <n-radio :value="false">按产品/SKU授权</n-radio>
 </n-space>
 </n-radio-group>
 </n-form-item>

 <template v-if="!formValue.isAllScope">
 <n-form-item label="授权范围">
 <n-tree-select
 v-model:value="formValue.scopeSelectedKeys"
 :options="scopeTreeOptions"
 multiple
 clearable
 checkable
 default-expand-all
 check-strategy="child"
 placeholder="请选择产品/SKU"
 class="scope-tree-select"
 />
 </n-form-item>
 </template>
 </n-form>
 </n-spin>
 </n-drawer-content>
 </n-drawer>
 </div>
</template>

<script lang="ts" setup>
import { computed, ref } from 'vue';
import { TreeSelectOption, useMessage } from 'naive-ui';
import { Edit, ScopeOptions, View } from '@/api/travelVerifyStaff';
import { ScopeOption, State, newState, rules } from './model';
import { adaModalWidth } from '@/utils/hotgo';

const emit = defineEmits(['reloadTable']);
const message = useMessage();
const formRef = ref<any>({});
const loading = ref(false);
const dialogWidth = computed(() => adaModalWidth(720));
const showModal = ref(false);
const formBtnLoading = ref(false);
const formValue = ref<State>(newState(null));
const scopeOptions = ref<ScopeOption[]>([]);

const scopeTreeOptions = computed<TreeSelectOption[]>(() =>
 scopeOptions.value.map((product) => {
   const productId = Number(product.id ||0);
   const productLabel = formatLabel(product.title, productId);
   const allSkuNode: TreeSelectOption = {
   label: '该产品全部SKU',
   key: toScopeKey(productId,0),
   value: toScopeKey(productId,0),
   isLeaf: true,
   };
   const skuNodes: TreeSelectOption[] = (product.skuList || []).map((sku) => {
   const skuId = Number(sku.id ||0);
   return {
   label: formatLabel(sku.name, skuId),
   key: toScopeKey(productId, skuId),
   value: toScopeKey(productId, skuId),
   isLeaf: true,
   };
   });
   return {
   label: productLabel,
   key: `product:${productId}`,
   value: `product:${productId}`,
   disabled: true,
   children: [...skuNodes],
   };
 }),
);

function formatLabel(label: string, id: number) {
 const text = `${label || ''}`.trim();
 return text || `#${id}`;
}

function toScopeKey(productId: number, skuId: number) {
 return `${Number(productId ||0)}:${Number(skuId ||0)}`;
}

function parseScopeKey(key: string) {
 const [productIdText, skuIdText] = `${key || ''}`.split(':');
 const productId = Number(productIdText);
 const skuId = Number(skuIdText);
 if (!Number.isInteger(productId) || productId <=0) {
 return null;
 }
 if (!Number.isInteger(skuId) || skuId <0) {
 return null;
 }
 return { productId, skuId };
}

function toScopeSelectedKeys(scopeItems: Array<{ productId: number; skuId: number }>) {
 const unique = new Set<string>();
 for (const item of scopeItems || []) {
 const productId = Number(item.productId ||0);
 const skuId = Number(item.skuId ||0);
 if (!Number.isInteger(productId) || productId <=0) {
 continue;
 }
 if (!Number.isInteger(skuId) || skuId <0) {
 continue;
 }
 unique.add(toScopeKey(productId, skuId));
 }
 return Array.from(unique);
}

function toScopeItemsFromSelectedKeys(selectedKeys: Array<string | number>) {
   const grouped = new Map<number, Set<number>>();
   for (const key of selectedKeys || []) {
   const parsed = parseScopeKey(String(key));
   if (!parsed) {
   continue;
   }
   if (!grouped.has(parsed.productId)) {
   grouped.set(parsed.productId, new Set<number>());
   }
   grouped.get(parsed.productId)?.add(parsed.skuId);
   }

   const result: Array<{ productId: number; skuId: number }> = [];
   for (const [productId, skuIds] of grouped.entries()) {
   if (skuIds.has(0)) {
   result.push({ productId, skuId:0 });
   continue;
   }
   for (const skuId of skuIds.values()) {
   result.push({ productId, skuId });
   }
   }
   return result;
}

function validateScopeItems() {
   if (formValue.value.isAllScope) {
   return true;
   }
   const scopeItems = toScopeItemsFromSelectedKeys(formValue.value.scopeSelectedKeys || []);
   if (scopeItems.length ===0) {
   message.error('请选择至少一条授权范围');
   return false;
   }
   return true;
}

function normalizePayload() {
 const payload = { ...formValue.value };
 if (payload.isAllScope) {
 payload.scopeItems = [];
 return payload;
 }
 payload.scopeItems = toScopeItemsFromSelectedKeys(payload.scopeSelectedKeys || []).map((item) => ({
 productId: Number(item.productId ||0),
 skuId: Number(item.skuId ||0),
 }));
 return payload;
}

function confirmForm(e) {
 e.preventDefault();
 formBtnLoading.value = true;
 formRef.value?.validate((errors) => {
 if (!errors) {
 if (!validateScopeItems()) {
 formBtnLoading.value = false;
 return;
 }
 Edit(normalizePayload())
 .then(() => {
 message.success('操作成功');
 formBtnLoading.value = false;
 closeForm();
 emit('reloadTable');
 })
 .catch(() => {
 message.error('操作失败，请稍后再试');
 formBtnLoading.value = false;
 });
 } else {
 message.error('请填写完整信息');
 formBtnLoading.value = false;
 }
 });
}

async function loadScopeOptions() {
 const res = await ScopeOptions();
 scopeOptions.value = res.list || [];
}

async function getInfo(id) {
 const res = await View({ id, isLanguage: true });
 const next = newState(res);
 next.password = '';
 next.scopeItems = next.scopeItems || [];
 next.scopeSelectedKeys = toScopeSelectedKeys(next.scopeItems);
 formValue.value = next;
}

async function openModal(state: State) {
 showModal.value = true;
 loading.value = true;

 await loadScopeOptions();

 if (!state || state.id <1) {
 const next = newState(state);
 next.isAllScope = true;
 next.scopeItems = [];
 next.scopeSelectedKeys = [];
 formValue.value = next;
 loading.value = false;
 return;
 }

 await getInfo(state.id);
 loading.value = false;
}

function closeForm() {
 showModal.value = false;
 loading.value = false;
}

defineExpose({
 openModal,
});
</script>

<style lang="less" scoped>
.level-detail-div {
 &-title {
 display: flex;
 align-items: center;
 font-weight:500;
 font-size:16px;
 color: #3d3d3d;
 line-height:22px;
 margin-bottom:15px;

 div {
 margin-right:5px;
 width:6px;
 height:15px;
 background: #053dc8;
 }
 }
}

.scope-tree-select {
 width:100%;
}
</style>
