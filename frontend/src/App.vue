<script setup lang="ts">
import { ref } from 'vue'
import PasswordVerifier from './components/PasswordVerifier.vue'
import PasswordGenerator from './components/PasswordGenerator.vue'
import { Shield, KeyRound, CheckCircle2 } from 'lucide-vue-next'

const activeTab = ref<'verifier' | 'generator'>('verifier')
</script>

<template>
  <main class="min-h-screen bg-slate-950 text-slate-100 flex flex-col items-center justify-center px-4 py-12">
    <header class="text-center mb-8 space-y-2">
      <div class="inline-flex items-center justify-center p-3 bg-emerald-500/10 border border-emerald-500/20 rounded-2xl mb-2">
        <Shield class="w-8 h-8 text-emerald-400" />
      </div>
      <h1 class="text-3xl font-bold tracking-tight text-slate-100">KeyForge</h1>
      <p class="text-slate-400 text-sm max-w-md mx-auto">
        Suite profesional de generación y validación de seguridad de contraseñas.
      </p>

      <!-- Selector principal de vistas -->
      <div class="flex justify-center gap-2 pt-4">
        <button
          @click="activeTab = 'verifier'"
          class="flex items-center gap-2 px-4 py-2 rounded-xl text-sm font-medium transition"
          :class="activeTab === 'verifier' ? 'bg-emerald-500/20 text-emerald-300 border border-emerald-500/30' : 'bg-slate-900 text-slate-400 border border-slate-800 hover:text-slate-200'"
        >
          <CheckCircle2 class="w-4 h-4" />
          Verificador
        </button>
        <button
          @click="activeTab = 'generator'"
          class="flex items-center gap-2 px-4 py-2 rounded-xl text-sm font-medium transition"
          :class="activeTab === 'generator' ? 'bg-emerald-500/20 text-emerald-300 border border-emerald-500/30' : 'bg-slate-900 text-slate-400 border border-slate-800 hover:text-slate-200'"
        >
          <KeyRound class="w-4 h-4" />
          Generador
        </button>
      </div>
    </header>

    <div class="w-full">
      <PasswordVerifier v-if="activeTab === 'verifier'" />
      <PasswordGenerator v-else />
    </div>
  </main>
</template>
