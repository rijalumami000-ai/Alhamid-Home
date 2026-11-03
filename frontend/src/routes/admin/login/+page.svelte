<script lang="ts">
  import { api, setToken } from '$lib/api';
  import { goto } from '$app/navigation';
  import { Lock, User, ArrowRight, Loader2, Sparkles, AlertCircle } from 'lucide-svelte';

  let username = $state('admin');
  let password = $state('alhamidku123');
  let loading = $state(false);
  let errorMsg = $state<string | null>(null);

  async function handleLogin(e: SubmitEvent) {
    e.preventDefault();
    loading = true;
    errorMsg = null;

    try {
      const res = await api.login({ username, password });
      setToken(res.token);
      goto('/admin');
    } catch (err: any) {
      errorMsg = err.message || 'Login gagal. Periksa username dan kata sandi Anda.';
    } finally {
      loading = false;
    }
  }
</script>

<svelte:head>
  <title>Login Administrator | CMS Al-Hamid</title>
</svelte:head>

<div class="min-h-screen bg-slate-950 flex items-center justify-center p-4 relative overflow-hidden">
  <!-- Subtle background glow -->
  <div class="absolute -top-40 -left-40 w-96 h-96 bg-emerald-600/20 rounded-full blur-3xl"></div>
  <div class="absolute -bottom-40 -right-40 w-96 h-96 bg-amber-500/10 rounded-full blur-3xl"></div>

  <div class="w-full max-w-md bg-white/95 backdrop-blur-xl rounded-3xl shadow-2xl border border-white/20 p-8 sm:p-10 space-y-8 relative z-10">
    <!-- Brand Emblem -->
    <div class="text-center space-y-3">
      <div class="w-14 h-14 rounded-2xl bg-gradient-to-tr from-emerald-900 to-emerald-600 text-amber-300 font-serif-luxury font-bold text-2xl flex items-center justify-center mx-auto shadow-lg shadow-emerald-950/30">
        AH
      </div>
      <div>
        <h1 class="font-serif-luxury font-bold text-2xl text-slate-900">
          Portal CMS Al-Hamid
        </h1>
        <p class="text-xs text-slate-500 mt-1">
          Sistem Pengelolaan Konten Website Resmi Pesantren
        </p>
      </div>
    </div>

    <!-- Error Alert -->
    {#if errorMsg}
      <div class="p-3.5 rounded-xl bg-red-50 border border-red-200 text-red-700 text-xs font-semibold flex items-center gap-2">
        <AlertCircle class="w-4 h-4 shrink-0" />
        <span>{errorMsg}</span>
      </div>
    {/if}

    <!-- Login Form -->
    <form onsubmit={handleLogin} class="space-y-4">
      <div class="space-y-1.5">
        <label for="username" class="text-xs font-bold text-slate-700 uppercase tracking-wide">
          Username atau Email
        </label>
        <div class="relative">
          <input
            id="username"
            type="text"
            required
            bind:value={username}
            placeholder="Username admin"
            class="w-full pl-10 pr-4 py-3 rounded-xl border border-slate-200 text-sm focus:outline-none focus:ring-2 focus:ring-emerald-600/30 font-medium"
          />
          <User class="w-4 h-4 text-slate-400 absolute left-3.5 top-3.5" />
        </div>
      </div>

      <div class="space-y-1.5">
        <label for="password" class="text-xs font-bold text-slate-700 uppercase tracking-wide">
          Kata Sandi
        </label>
        <div class="relative">
          <input
            id="password"
            type="password"
            required
            bind:value={password}
            placeholder="••••••••"
            class="w-full pl-10 pr-4 py-3 rounded-xl border border-slate-200 text-sm focus:outline-none focus:ring-2 focus:ring-emerald-600/30 font-medium"
          />
          <Lock class="w-4 h-4 text-slate-400 absolute left-3.5 top-3.5" />
        </div>
      </div>

      <button
        type="submit"
        disabled={loading}
        class="w-full py-3.5 rounded-xl bg-gradient-to-r from-emerald-800 to-emerald-700 hover:from-emerald-700 hover:to-emerald-600 text-white font-bold text-sm shadow-lg shadow-emerald-950/20 transition duration-200 flex items-center justify-center gap-2 disabled:opacity-50"
      >
        {#if loading}
          <Loader2 class="w-4 h-4 animate-spin" />
          <span>Memverifikasi...</span>
        {:else}
          <span>Masuk ke Dashboard</span>
          <ArrowRight class="w-4 h-4" />
        {/if}
      </button>
    </form>

    <div class="pt-2 text-center text-xs text-slate-400">
      <a href="/" class="hover:text-emerald-700 transition font-medium">
        &larr; Kembali ke Website Publik
      </a>
    </div>
  </div>
</div>

