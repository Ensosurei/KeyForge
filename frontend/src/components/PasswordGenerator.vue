<script setup lang="ts">
import { ref, reactive, onMounted } from 'vue'
import { generatePassword, generatePassphrase, type GeneratorOptions, type PassphraseOptions } from '../services/api'
import { KeyRound, Copy, Check, RefreshCw } from 'lucide-vue-next'

const mode = ref<'password' | 'passphrase'>('password')
const generatedResult = ref('')
const copied = ref(false)

const passwordOpts = reactive<GeneratorOptions>({
  length: 16,
  includeUpper: true,
  includeLower: true,
  includeDigits: true,
  includeSymbols: true,
})

const passphraseOpts = reactive<PassphraseOptions>({
  wordCount: 4,
  separator: '-',
  capitalize: true,
  enableLeet: false,
})

const handleGenerate = async () => {
  if (mode.value === 'password') {
    generatedResult.value = await generatePassword(passwordOpts)
  } else {
    generatedResult.value = await generatePassphrase(passphraseOpts)
  }
}

const copyToClipboard = async () => {
  if (!generatedResult.value) return
  await navigator.clipboard.writeText(generatedResult.value)
  copied.value = true
  setTimeout(() => {
    copied.value = false
  }, 2000)
}

onMounted(() => {
  handleGenerate()
})
</script>

<template>
  <div class="max-w-xl mx-auto bg-slate-900 border border-slate-800 rounded-2xl p-6 shadow-2xl space-y-6">
    <div class="space-y-2">
      <h2 class="text-xl font-semibold text-slate-100 flex items-center gap-2">
        <KeyRound class="w-6 h-6 text-emerald-400" />
        Generador KeyForge
      </h2>
      <p class="text-sm text-slate-400">
        Crea contraseñas criptográficamente seguras o frases secretas fáciles de recordar.
      </p>
    </div>

    <!-- Pestañas para Cambiar de Modo -->
    <div class="flex bg-slate-950 p-1 rounded-xl border border-slate-800">
      <button
        @click="mode = 'password'; handleGenerate()"
        class="flex-1 py-2 text-sm font-medium rounded-lg transition cursor-pointer"
        :class="mode === 'password' ? 'bg-slate-800 text-emerald-400 shadow' : 'text-slate-400 hover:text-slate-200'"
      >
        Contraseña Aleatoria
      </button>
      <button
        @click="mode = 'passphrase'; handleGenerate()"
        class="flex-1 py-2 text-sm font-medium rounded-lg transition cursor-pointer"
        :class="mode === 'passphrase' ? 'bg-slate-800 text-emerald-400 shadow' : 'text-slate-400 hover:text-slate-200'"
      >
        Frase Secreta
      </button>
    </div>

    <!-- Resultado Generado + Botones de Acción -->
    <div class="relative flex items-center">
      <input
        type="text"
        readonly
        :value="generatedResult"
        class="w-full bg-slate-950 border border-slate-800 rounded-xl px-4 py-3 text-slate-100 font-mono text-lg tracking-wide focus:outline-none pr-24"
      />
      <div class="absolute right-2 flex items-center gap-1">
        <button
          @click="handleGenerate"
          title="Regenerar"
          class="p-2 text-slate-400 hover:text-emerald-400 hover:bg-slate-800 rounded-lg transition cursor-pointer"
        >
          <RefreshCw class="w-5 h-5" />
        </button>
        <button
          @click="copyToClipboard"
          title="Copiar"
          class="p-2 text-slate-400 hover:text-emerald-400 hover:bg-slate-800 rounded-lg transition cursor-pointer relative"
        >
          <Check v-if="copied" class="w-5 h-5 text-emerald-400" />
          <Copy v-else class="w-5 h-5" />
        </button>
      </div>
    </div>

    <!-- Controles para Contraseña Aleatoria -->
    <div v-if="mode === 'password'" class="space-y-4 pt-2">
      <div class="space-y-2">
        <div class="flex justify-between text-sm">
          <span class="text-slate-300">Longitud:</span>
          <span class="font-mono text-emerald-400 font-semibold">{{ passwordOpts.length }} caracteres</span>
        </div>
        <input
          type="range"
          min="8"
          max="64"
          v-model.number="passwordOpts.length"
          @input="handleGenerate"
          class="w-full accent-emerald-500 bg-slate-950 rounded-lg h-2 cursor-pointer"
        />
      </div>

      <div class="grid grid-cols-2 gap-3 text-sm">
        <label class="flex items-center gap-2 text-slate-300 cursor-pointer">
          <input
            type="checkbox"
            v-model="passwordOpts.includeUpper"
            @change="handleGenerate"
            class="accent-emerald-500 rounded cursor-pointer"
          />
          Mayúsculas (A-Z)
        </label>
        <label class="flex items-center gap-2 text-slate-300 cursor-pointer">
          <input
            type="checkbox"
            v-model="passwordOpts.includeLower"
            @change="handleGenerate"
            class="accent-emerald-500 rounded cursor-pointer"
          />
          Minúsculas (a-z)
        </label>
        <label class="flex items-center gap-2 text-slate-300 cursor-pointer">
          <input
            type="checkbox"
            v-model="passwordOpts.includeDigits"
            @change="handleGenerate"
            class="accent-emerald-500 rounded cursor-pointer"
          />
          Números (0-9)
        </label>
        <label class="flex items-center gap-2 text-slate-300 cursor-pointer">
          <input
            type="checkbox"
            v-model="passwordOpts.includeSymbols"
            @change="handleGenerate"
            class="accent-emerald-500 rounded cursor-pointer"
          />
          Símbolos (!@#$)
        </label>
      </div>
    </div>

    <!-- Controles para Frase Secreta -->
    <div v-else class="space-y-4 pt-2">
      <div class="space-y-2">
        <div class="flex justify-between text-sm">
          <span class="text-slate-300">Número de Palabras:</span>
          <span class="font-mono text-emerald-400 font-semibold">{{ passphraseOpts.wordCount }} palabras</span>
        </div>
        <input
          type="range"
          min="3"
          max="8"
          v-model.number="passphraseOpts.wordCount"
          @input="handleGenerate"
          class="w-full accent-emerald-500 bg-slate-950 rounded-lg h-2 cursor-pointer"
        />
      </div>

      <div class="flex items-center justify-between text-sm">
        <span class="text-slate-300">Separador:</span>
        <select
          v-model="passphraseOpts.separator"
          @change="handleGenerate"
          class="bg-slate-950 border border-slate-800 rounded-lg px-3 py-1.5 text-slate-200 font-mono text-sm focus:outline-none focus:border-emerald-500 cursor-pointer"
        >
          <option value="-">Guion (-)</option>
          <option value="_">Guion Bajo (_)</option>
          <option value=".">Punto (.)</option>
          <option value=" ">Espacio (" ")</option>
        </select>
      </div>

      <!-- Casillas de verificación para Mayúscula Inicial y Leet Speak -->
      <div class="grid grid-cols-2 gap-3 text-sm pt-2">
        <label class="flex items-center gap-2 text-slate-300 cursor-pointer">
          <input
            type="checkbox"
            v-model="passphraseOpts.capitalize"
            @change="handleGenerate"
            class="accent-emerald-500 rounded cursor-pointer"
          />
          Mayúscula Inicial
        </label>
        <label class="flex items-center gap-2 text-slate-300 cursor-pointer">
          <input
            type="checkbox"
            v-model="passphraseOpts.enableLeet"
            @change="handleGenerate"
            class="accent-emerald-500 rounded cursor-pointer"
          />
          Leet Speak (fu3g0)
        </label>
      </div>
    </div>
  </div>
</template>
