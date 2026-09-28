<script setup lang="ts">
import { ref, watch } from 'vue'
import { verifyPassword, type VerificationResponse } from '../services/api'
import { CheckCircle2, XCircle, ShieldCheck, ShieldAlert, Eye, EyeOff } from 'lucide-vue-next'

const password = ref('')
const showPassword = ref(false)
const result = ref<VerificationResponse | null>(null)

// Evaluar la contraseña en tiempo real conforme el usuario escribe
watch(password, async (newVal) => {
  if (newVal.length === 0) {
    result.value = null
    return
  }
  result.value = await verifyPassword(newVal)
})

const toggleShowPassword = () => {
  showPassword.value = !showPassword.value
}
</script>

<template>
  <div class="max-w-xl mx-auto bg-slate-900 border border-slate-800 rounded-2xl p-6 shadow-2xl space-y-6">
    <div class="space-y-2">
      <h2 class="text-xl font-semibold text-slate-100 flex items-center gap-2">
        <ShieldCheck class="w-6 h-6 text-emerald-400" />
        Verificador de Seguridad
      </h2>
      <p class="text-sm text-slate-400">
        Escribe una contraseña para analizar su nivel de fortaleza y reglas de seguridad.
      </p>
    </div>

    <!-- Campo de Entrada -->
    <div class="relative">
      <input
        :type="showPassword ? 'text' : 'password'"
        v-model="password"
        placeholder="Ingresa tu contraseña aquí..."
        class="w-full bg-slate-950 border border-slate-800 rounded-xl px-4 py-3 text-slate-100 placeholder-slate-500 focus:outline-none focus:border-emerald-500 transition pr-12 font-mono"
      />
      <button
        @click="toggleShowPassword"
        type="button"
        class="absolute right-3 top-1/2 -translate-y-1/2 text-slate-400 hover:text-slate-200 transition"
      >
        <EyeOff v-if="showPassword" class="w-5 h-5" />
        <Eye v-else class="w-5 h-5" />
      </button>
    </div>

    <!-- Barra de Fortaleza y Estado -->
    <div v-if="result" class="space-y-4">
      <div class="flex justify-between items-center text-sm">
        <span class="text-slate-400">Fortaleza:</span>
        <span
          class="font-semibold"
          :class="{
            'text-red-400': result.score < 50,
            'text-yellow-400': result.score >= 50 && result.score < 80,
            'text-emerald-400': result.score >= 80
          }"
        >
          {{ result.strength }} ({{ result.score }}%)
        </span>
      </div>

      <!-- Barra de Progreso -->
      <div class="w-full h-2 bg-slate-950 rounded-full overflow-hidden">
        <div
          class="h-full transition-all duration-300"
          :class="{
            'bg-red-500': result.score < 50,
            'bg-yellow-500': result.score >= 50 && result.score < 80,
            'bg-emerald-500': result.score >= 80
          }"
          :style="{ width: `${result.score}%` }"
        ></div>
      </div>

      <!-- Lista de Checks Dinámicos -->
      <div class="space-y-3 pt-2">
        <div
          v-for="check in result.checks"
          :key="check.id"
          class="flex items-start gap-3 p-3 rounded-lg border transition"
          :class="check.passed ? 'bg-emerald-950/20 border-emerald-900/50 text-emerald-300' : 'bg-slate-950 border-slate-800 text-slate-400'"
        >
          <CheckCircle2 v-if="check.passed" class="w-5 h-5 text-emerald-400 shrink-0 mt-0.5" />
          <XCircle v-else class="w-5 h-5 text-slate-600 shrink-0 mt-0.5" />

          <div class="space-y-0.5">
            <p class="text-sm font-medium" :class="check.passed ? 'text-emerald-200' : 'text-slate-300'">
              {{ check.label }}
            </p>
            <p class="text-xs text-slate-400">
              {{ check.description }}
            </p>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
