export interface CheckResult {
  id: string;
  label: string;
  passed: boolean;
  description: string;
}

export interface VerificationResponse {
  score: number;
  strength: string;
  isSecure: boolean;
  checks: CheckResult[];
  feedback: string[];
}

export interface GeneratorOptions {
  length: number;
  includeUpper: boolean;
  includeLower: boolean;
  includeDigits: boolean;
  includeSymbols: boolean;
}

export interface PassphraseOptions {
  wordCount: number;
  separator: string;
  capitalize: boolean;
  enableLeet: boolean;
}

export async function verifyPassword(password: string): Promise<VerificationResponse> {
  try {
    const response = await fetch('/api/verify', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ password }),
    });

    if (!response.ok) {
      throw new Error('Error al conectar con la API');
    }

    return await response.json();
  } catch {
    return evaluateLocal(password);
  }
}

export async function generatePassword(opts: GeneratorOptions): Promise<string> {
  try {
    const res = await fetch('/api/generate', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(opts),
    });
    const data = await res.json();
    return data.password;
  } catch {
    return generateLocalPassword(opts);
  }
}

export async function generatePassphrase(opts: PassphraseOptions): Promise<string> {
  try {
    const res = await fetch('/api/passphrase', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(opts),
    });
    const data = await res.json();
    return data.passphrase;
  } catch {
    return generateLocalPassphrase(opts);
  }
}

function evaluateLocal(password: string): VerificationResponse {
  const hasLength = password.length >= 12;
  const hasUpper = /[A-Z]/.test(password);
  const hasLower = /[a-z]/.test(password);
  const hasDigit = /[0-9]/.test(password);
  const hasSymbol = /[^A-Za-z0-9]/.test(password);
  const hasVariety = hasUpper && hasLower && hasDigit && hasSymbol;
  const isObvious = /123456|password|qwerty|contraseña/i.test(password);

  const checks: CheckResult[] = [
    {
      id: 'length',
      label: 'Longitud suficiente',
      passed: hasLength,
      description: 'La contraseña debe contener al menos 12 caracteres.',
    },
    {
      id: 'variety',
      label: 'Variedad de caracteres',
      passed: hasVariety,
      description: 'Debe combinar letras mayúsculas, minúsculas, números y símbolos.',
    },
    {
      id: 'avoid_obvious',
      label: 'Evita patrones obvios',
      passed: !isObvious,
      description: 'No debe contener palabras comunes ni secuencias sencillas.',
    },
  ];

  const passedCount = checks.filter((c) => c.passed).length;
  const score = password.length > 0 ? Math.round((passedCount / checks.length) * 100) : 0;

  return {
    score,
    strength: score >= 90 ? 'Muy Fuerte' : score >= 60 ? 'Fuerte' : 'Débil',
    isSecure: hasLength && hasVariety && !isObvious,
    checks,
    feedback: [],
  };
}

function generateLocalPassword(opts: GeneratorOptions): string {
  let charset = '';
  if (opts.includeLower) charset += 'abcdefghijklmnopqrstuvwxyz';
  if (opts.includeUpper) charset += 'ABCDEFGHIJKLMNOPQRSTUVWXYZ';
  if (opts.includeDigits) charset += '0123456789';
  if (opts.includeSymbols) charset += '!@#$%^&*()_+-=[]{}|;:,.<>?';

  if (!charset) charset = 'abcdefghijklmnopqrstuvwxyz0123456789';

  let res = '';
  for (let i = 0; i < opts.length; i++) {
    res += charset.charAt(Math.floor(Math.random() * charset.length));
  }
  return res;
}

function generateLocalPassphrase(opts: PassphraseOptions): string {
  const words = [
    'fuego', 'codigo', 'llave', 'fuerza', 'sombra', 'trueno', 'dragon',
    'viento', 'tormenta', 'cristal', 'guardian', 'secreto', 'vortex', 'orbita'
  ];

  const selected: string[] = [];
  let last = '';

  while (selected.length < opts.wordCount) {
    let word = words[Math.floor(Math.random() * words.length)];
    if (word === last) continue;
    last = word;

    if (opts.capitalize) {
      word = word.charAt(0).toUpperCase() + word.slice(1);
    }

    if (opts.enableLeet) {
      word = word
        .replace(/a/gi, '4')
        .replace(/e/gi, '3')
        .replace(/i/gi, '1')
        .replace(/o/gi, '0')
        .replace(/s/gi, '5');
    }

    selected.push(word);
  }

  return selected.join(opts.separator);
}
