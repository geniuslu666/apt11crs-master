<template>
  <div class="flex min-h-svh flex-col items-center justify-center bg-muted p-6 md:p-10">
    <div class="flex w-full max-w-sm flex-col gap-6 md:max-w-4xl">

      <!-- Card (overflow-hidden + p-0, matching shadcn Card) -->
      <div class="overflow-hidden rounded-xl border border-border bg-background shadow-sm">
        <div class="grid p-0 md:grid-cols-2">

          <!-- Left: Form (FieldGroup = flex flex-col gap-6) -->
          <form class="flex flex-col gap-6 p-6 md:p-8" @submit.prevent="handleLogin">

            <!-- Header (centered) -->
            <div class="flex flex-col items-center gap-2 text-center">
              <h1 class="text-2xl font-bold">欢迎回来</h1>
              <p class="text-balance text-sm text-muted-foreground">
                登录您的住一CRS账号
              </p>
            </div>

            <!-- Username (Field = flex flex-col gap-2) -->
            <div class="flex flex-col gap-2">
              <label for="username" class="text-sm font-medium leading-none">
                用户名
              </label>
              <input
                id="username"
                v-model="form.username"
                type="text"
                placeholder="请输入用户名"
                autocomplete="username"
                @keyup.enter="handleLogin"
                :class="[
                  'flex h-9 w-full rounded-md border bg-transparent px-3 py-1 text-sm shadow-sm transition-colors',
                  'placeholder:text-muted-foreground focus-visible:outline-none focus-visible:ring-1',
                  'disabled:cursor-not-allowed disabled:opacity-50',
                  errors.username ? 'border-destructive focus-visible:ring-destructive' : 'border-input focus-visible:ring-ring'
                ]"
              />
              <p v-if="errors.username" class="text-xs text-destructive">{{ errors.username }}</p>
            </div>

            <!-- Password (Field) -->
            <div class="flex flex-col gap-2">
              <div class="flex items-center">
                <label for="password" class="text-sm font-medium leading-none">密码</label>
                <button
                  type="button"
                  @click="handleResetPassword"
                  class="ml-auto text-sm underline-offset-4 hover:underline"
                >
                  忘记密码？
                </button>
              </div>
              <div class="relative">
                <input
                  id="password"
                  v-model="form.pass"
                  :type="showPassword ? 'text' : 'password'"
                  placeholder="请输入密码"
                  autocomplete="current-password"
                  @keyup.enter="handleLogin"
                  :class="[
                    'flex h-9 w-full rounded-md border bg-transparent px-3 py-1 pr-10 text-sm shadow-sm transition-colors',
                    'placeholder:text-muted-foreground focus-visible:outline-none focus-visible:ring-1',
                    'disabled:cursor-not-allowed disabled:opacity-50',
                    errors.pass ? 'border-destructive focus-visible:ring-destructive' : 'border-input focus-visible:ring-ring'
                  ]"
                />
                <button
                  type="button"
                  tabindex="-1"
                  @click="showPassword = !showPassword"
                  class="absolute right-3 top-1/2 -translate-y-1/2 text-muted-foreground hover:text-foreground transition-colors"
                >
                  <!-- Eye open -->
                  <svg v-if="showPassword" xmlns="http://www.w3.org/2000/svg" width="15" height="15" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M2 12s3-7 10-7 10 7 10 7-3 7-10 7-10-7-10-7Z"/><circle cx="12" cy="12" r="3"/></svg>
                  <!-- Eye off -->
                  <svg v-else xmlns="http://www.w3.org/2000/svg" width="15" height="15" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M9.88 9.88a3 3 0 1 0 4.24 4.24"/><path d="M10.73 5.08A10.43 10.43 0 0 1 12 5c7 0 10 7 10 7a13.16 13.16 0 0 1-1.67 2.68"/><path d="M6.61 6.61A13.526 13.526 0 0 0 2 12s3 7 10 7a9.74 9.74 0 0 0 5.39-1.61"/><line x1="2" x2="22" y1="2" y2="22"/></svg>
                </button>
              </div>
              <p v-if="errors.pass" class="text-xs text-destructive">{{ errors.pass }}</p>
            </div>

            <!-- Captcha (conditional Field) -->
            <div v-if="codeBase64" class="flex flex-col gap-2">
              <label class="text-sm font-medium leading-none">验证码</label>
              <div class="flex gap-2">
                <input
                  v-model="form.code"
                  type="text"
                  placeholder="请输入验证码"
                  @keyup.enter="handleLogin"
                  class="flex h-9 flex-1 min-w-0 rounded-md border border-input bg-transparent px-3 py-1 text-sm shadow-sm transition-colors placeholder:text-muted-foreground focus-visible:outline-none focus-visible:ring-1 focus-visible:ring-ring"
                />
                <img
                  :src="codeBase64"
                  @click="refreshCode"
                  title="点击刷新"
                  class="h-9 w-28 flex-shrink-0 cursor-pointer rounded-md border border-border object-cover hover:opacity-80 transition-opacity"
                  alt="验证码"
                />
              </div>
            </div>

            <!-- Remember me -->
            <div class="flex items-center gap-2">
              <input
                type="checkbox"
                id="remember"
                v-model="autoLogin"
                class="h-4 w-4 cursor-pointer rounded border-input accent-primary"
              />
              <label for="remember" class="cursor-pointer select-none text-sm text-muted-foreground">
                记住登录状态
              </label>
            </div>

            <!-- Submit button (Field + Button default variant) -->
            <button
              type="submit"
              :disabled="loading"
              class="flex h-9 w-full items-center justify-center gap-2 rounded-md bg-foreground px-4 py-2 text-sm font-medium text-background shadow hover:bg-foreground/90 focus-visible:outline-none focus-visible:ring-1 focus-visible:ring-ring disabled:pointer-events-none disabled:opacity-50 transition-colors"
            >
              <svg v-if="loading" class="h-4 w-4 animate-spin" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24">
                <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"/>
                <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"/>
              </svg>
              {{ loading ? '登录中...' : '登录' }}
            </button>

            <!-- FieldSeparator: "Or continue with" -->
            <div class="relative text-center text-xs text-muted-foreground after:absolute after:inset-0 after:top-1/2 after:z-0 after:flex after:items-center after:border-t after:border-border">
              <span class="relative z-10 bg-background px-2">或通过以下方式</span>
            </div>

            <!-- Social buttons grid (Field className="grid grid-cols-3 gap-4") -->
            <div class="grid grid-cols-3 gap-4">
              <!-- Apple -->
              <button
                type="button"
                class="flex h-9 w-full items-center justify-center rounded-md border border-input bg-background shadow-sm hover:bg-accent hover:text-accent-foreground focus-visible:outline-none focus-visible:ring-1 focus-visible:ring-ring transition-colors"
              >
                <svg xmlns="http://www.w3.org/2000/svg" class="h-4 w-4" viewBox="0 0 24 24" fill="currentColor">
                  <path d="M12.152 6.896c-.948 0-2.415-1.078-3.96-1.04-2.04.027-3.91 1.183-4.961 3.014-2.117 3.675-.546 9.103 1.519 12.09 1.013 1.454 2.208 3.09 3.792 3.039 1.52-.065 2.09-.987 3.935-.987 1.831 0 2.35.987 3.96.948 1.637-.026 2.676-1.48 3.676-2.948 1.156-1.688 1.636-3.325 1.662-3.415-.039-.013-3.182-1.221-3.22-4.857-.026-3.04 2.48-4.494 2.597-4.559-1.429-2.09-3.623-2.324-4.39-2.376-2-.156-3.675 1.09-4.61 1.09zM15.53 3.83c.843-1.012 1.4-2.427 1.245-3.83-1.207.052-2.662.805-3.532 1.818-.78.896-1.454 2.338-1.273 3.714 1.338.104 2.715-.688 3.559-1.701"/>
                </svg>
                <span class="sr-only">Apple 登录</span>
              </button>
              <!-- Google -->
              <button
                type="button"
                class="flex h-9 w-full items-center justify-center rounded-md border border-input bg-background shadow-sm hover:bg-accent hover:text-accent-foreground focus-visible:outline-none focus-visible:ring-1 focus-visible:ring-ring transition-colors"
              >
                <svg xmlns="http://www.w3.org/2000/svg" class="h-4 w-4" viewBox="0 0 24 24" fill="currentColor">
                  <path d="M12.48 10.92v3.28h7.84c-.24 1.84-.853 3.187-1.787 4.133-1.147 1.147-2.933 2.4-6.053 2.4-4.827 0-8.6-3.893-8.6-8.72s3.773-8.72 8.6-8.72c2.6 0 4.507 1.027 5.907 2.347l2.307-2.307C18.747 1.44 16.133 0 12.48 0 5.867 0 .307 5.387.307 12s5.56 12 12.173 12c3.573 0 6.267-1.173 8.373-3.36 2.16-2.16 2.84-5.213 2.84-7.667 0-.76-.053-1.467-.173-2.053H12.48z"/>
                </svg>
                <span class="sr-only">Google 登录</span>
              </button>
              <!-- Meta -->
              <button
                type="button"
                class="flex h-9 w-full items-center justify-center rounded-md border border-input bg-background shadow-sm hover:bg-accent hover:text-accent-foreground focus-visible:outline-none focus-visible:ring-1 focus-visible:ring-ring transition-colors"
              >
                <svg xmlns="http://www.w3.org/2000/svg" class="h-4 w-4" viewBox="0 0 24 24" fill="currentColor">
                  <path d="M6.915 4.03c-1.968 0-3.683 1.28-4.871 3.113C.704 9.208 0 11.883 0 14.449c0 .706.07 1.369.21 1.973a6.624 6.624 0 0 0 .265.86 5.297 5.297 0 0 0 .371.761c.696 1.159 1.818 1.927 3.593 1.927 1.497 0 2.633-.671 3.965-2.444.76-1.012 1.144-1.626 2.663-4.32l.756-1.339.186-.325c.061.1.121.196.183.3l2.152 3.595c.724 1.21 1.665 2.556 2.47 3.314 1.046.987 1.992 1.22 3.06 1.22 1.075 0 1.876-.355 2.455-.843a3.743 3.743 0 0 0 .81-.973c.542-.939.861-2.127.861-3.745 0-2.72-.681-5.357-2.084-7.45-1.282-1.912-2.957-2.93-4.716-2.93-1.047 0-2.088.467-3.053 1.308-.652.57-1.257 1.29-1.82 2.05-.69-.875-1.335-1.547-1.958-2.056-1.182-.966-2.315-1.303-3.454-1.303zm10.16 2.053c1.147 0 2.188.758 2.992 1.999 1.132 1.748 1.647 4.195 1.647 6.4 0 1.548-.368 2.9-1.839 2.9-.58 0-1.027-.23-1.664-1.004-.496-.601-1.343-1.878-2.832-4.358l-.617-1.028a44.908 44.908 0 0 0-1.255-1.98c.07-.109.141-.224.211-.327 1.12-1.667 2.118-2.602 3.358-2.602zm-10.201.553c1.265 0 2.058.791 2.675 1.446.307.327.737.871 1.234 1.579l-1.02 1.566c-.757 1.163-1.882 3.017-2.837 4.338-1.191 1.649-1.81 1.817-2.486 1.817-.524 0-1.038-.237-1.383-.794-.263-.426-.464-1.13-.464-2.046 0-2.221.63-4.535 1.66-6.088.454-.687.964-1.226 1.533-1.533a2.264 2.264 0 0 1 1.088-.285z"/>
                </svg>
                <span class="sr-only">Meta 登录</span>
              </button>
            </div>

            <!-- "Don't have an account?" (FieldDescription) -->
            <p class="text-center text-sm text-muted-foreground [&_a]:underline [&_a]:underline-offset-4 hover:[&_a]:text-foreground">
              没有账号？<a href="#">联系管理员</a>
            </p>

          </form>

          <!-- Right: Image panel -->
          <div class="relative hidden bg-muted md:block overflow-hidden">
            <!-- Brand gradient backdrop -->
            <div class="absolute inset-0 bg-gradient-to-br from-[#0a2540] via-[#0f2d56] to-[#0a1e35]"></div>
            <!-- Purple radial glow -->
            <div
              class="absolute inset-0"
              style="background: radial-gradient(ellipse at 35% 35%, rgba(99,91,255,0.3) 0%, transparent 60%), radial-gradient(ellipse at 75% 75%, rgba(99,91,255,0.12) 0%, transparent 50%);"
            ></div>
            <!-- Illustration overlay -->
            <img
              src="../../assets/images/new/login_bg.png"
              alt=""
              class="absolute inset-0 h-full w-full object-cover opacity-[0.07] pointer-events-none select-none"
            />
            <!-- Brand content -->
            <div class="absolute inset-0 flex flex-col items-center justify-center gap-5 p-10">
              <div class="flex h-14 w-14 items-center justify-center rounded-2xl bg-white/10 ring-1 ring-white/20 backdrop-blur-sm overflow-hidden shadow-lg">
                <img src="../../assets/images/new/login_logo.png" alt="logo" class="h-9 w-9 object-contain" />
              </div>
              <div class="text-center">
                <p class="text-[26px] font-bold text-white tracking-tight leading-tight">住一CRS</p>
                <p class="mt-2 text-sm text-blue-200/50 tracking-wide">业务中台</p>
              </div>
              <div class="mt-4 flex flex-col items-center gap-2 text-blue-200/40 text-xs">
                <span>统一管理 · 高效协作 · 实时数据</span>
              </div>
            </div>
          </div>

        </div>
      </div>

      <!-- Terms footer (FieldDescription below card) -->
      <p class="text-balance text-center text-xs text-muted-foreground [&_a]:underline [&_a]:underline-offset-4 hover:[&_a]:text-primary">
        点击登录即表示您同意我们的<a href="#">服务条款</a>和<a href="#">隐私政策</a>。
      </p>

    </div>
  </div>
</template>

<script lang="ts" setup>
import { reactive, ref, onMounted } from 'vue';
import { useMessage } from 'naive-ui';
import { useUserStore } from '@/store/modules/user';
import { ResultEnum } from '@/enums/httpEnum';
import { GetCaptcha } from '@/api/base';
import { useRoute, useRouter } from 'vue-router';
import { PageEnum } from '@/enums/pageEnum';
import { aesEcb } from '@/utils/encrypt';

const message = useMessage();
const userStore = useUserStore();
const loading = ref(false);
const codeBase64 = ref('');
const router = useRouter();
const route = useRoute();
const LOGIN_NAME = PageEnum.BASE_LOGIN_NAME;
const autoLogin = ref(true);
const showPassword = ref(false);

const form = reactive({
  username: '',
  pass: '',
  cid: '',
  code: '',
});

const errors = reactive({
  username: '',
  pass: '',
});

function validate(): boolean {
  errors.username = '';
  errors.pass = '';
  let valid = true;
  if (!form.username.trim()) { errors.username = '请输入用户名'; valid = false; }
  if (!form.pass.trim()) { errors.pass = '请输入密码'; valid = false; }
  return valid;
}

function handleResetPassword() {
  message.info('如果你忘记了密码，请联系管理员找回');
}

async function refreshCode() {
  if (userStore.loginConfig?.loginCaptchaSwitch !== 1) return;
  const data = await GetCaptcha();
  codeBase64.value = data.base64;
  form.cid = data.cid;
  form.code = '';
}

async function handleLoginResp(request: Promise<any>) {
  message.loading('登录中...');
  loading.value = true;
  try {
    const { code, message: msg } = await request;
    message.destroyAll();
    if (code == ResultEnum.SUCCESS) {
      const toPath = decodeURIComponent((route.query?.redirect || '/') as string);
      message.success('登录成功，即将进入系统');
      if (route.name === LOGIN_NAME) {
        await router.replace('/');
      } else {
        await router.replace(toPath);
      }
    } else {
      message.destroyAll();
      message.info(msg || '登录失败');
      await refreshCode();
    }
  } finally {
    loading.value = false;
  }
}

function handleLogin() {
  if (!validate()) return;
  if (userStore.loginConfig?.loginCaptchaSwitch === 1 && !form.code) {
    message.error('请输入验证码');
    return;
  }
  handleLoginResp(
    userStore.login({
      username: form.username,
      password: aesEcb.encrypt(form.pass),
      cid: form.cid,
      code: form.code,
    })
  );
}

onMounted(() => {
  setTimeout(refreshCode);
});
</script>
