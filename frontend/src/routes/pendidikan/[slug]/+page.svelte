<script lang="ts">
  import { onMount } from 'svelte';
  import { page } from '$app/state';
  import { api } from '$lib/api';
  import type { EducationUnit } from '$lib/types';
  import {
    GraduationCap,
    BookOpen,
    CheckCircle2,
    Calendar,
    Users,
    Building2,
    ClipboardList,
    ArrowRight,
    Sparkles,
    ShieldCheck
  } from 'lucide-svelte';

  let unit = $state<EducationUnit | null>(null);
  let loading = $state(true);
  let errorMsg = $state<string | null>(null);

  async function loadUnit(slug: string) {
    loading = true;
    errorMsg = null;
    try {
      unit = await api.getEducationUnitBySlug(slug);
    } catch (e: any) {
      errorMsg = e.message || 'Laman unit pendidikan tidak ditemukan';
    } finally {
      loading = false;
    }
  }

  $effect(() => {
    const slug = page.params.slug;
    if (slug) {
      loadUnit(slug);
    }
  });
</script>

<svelte:head>
  <title>{unit?.name ? `${unit.name} | Pondok Pesantren Al-Hamid` : 'Laman Pendidikan Al-Hamid'}</title>
</svelte:head>

{#if loading}
  <div class="py-32 text-center text-slate-400">
    <div class="animate-spin w-8 h-8 border-4 border-emerald-600 border-t-transparent rounded-full mx-auto mb-4"></div>
    Memuat informasi unit pendidikan...
  </div>
{:else if errorMsg || !unit}
  <div class="py-32 text-center max-w-md mx-auto px-4 space-y-4">
    <h1 class="text-2xl font-bold text-slate-800">Unit Pendidikan Tidak Ditemukan</h1>
    <p class="text-sm text-slate-500">Laman yang Anda tuju mungkin telah dipindahkan atau belum tersedia.</p>
    <a href="/pendidikan" class="inline-flex px-6 py-2.5 rounded-xl bg-emerald-800 text-white font-semibold text-sm">
      Kembali ke Daftar Pendidikan
    </a>
  </div>
{:else}
  <!-- Hero Section -->
  <section class="relative min-h-[420px] flex items-center bg-slate-950 text-white overflow-hidden">
    <div class="absolute inset-0">
      <img
        src={unit.hero_image}
        alt={unit.name}
        class="w-full h-full object-cover"
      />
      <div class="absolute inset-0 bg-gradient-to-r from-slate-950/95 via-slate-950/80 to-transparent"></div>
      <div class="absolute inset-0 bg-emerald-950/40 mix-blend-multiply"></div>
    </div>

    <div class="relative z-10 max-w-6xl mx-auto px-4 sm:px-8 py-20 w-full space-y-4">
      <div class="flex flex-wrap items-center gap-3">
        <span class="px-3 py-1 rounded-full text-xs font-bold uppercase tracking-wider shadow-sm {unit.type === 'formal' ? 'bg-blue-600 text-white' : 'bg-emerald-700 text-white'}">
          {unit.type === 'formal' ? 'Pendidikan Formal' : 'Pendidikan Non-Formal'}
        </span>
        {#if unit.badge}
          <span class="px-3 py-1 rounded-full bg-amber-400/20 text-amber-300 border border-amber-400/30 text-xs font-semibold backdrop-blur-md">
            {unit.badge}
          </span>
        {/if}
      </div>

      <h1 class="font-serif-luxury font-bold text-3xl sm:text-5xl lg:text-6xl text-white tracking-tight">
        {unit.name}
      </h1>

      <p class="text-base sm:text-lg text-emerald-100/90 max-w-2xl">
        {unit.tagline}
      </p>

      <div class="pt-4 flex items-center gap-4">
        <a
          href="/kontak"
          class="px-7 py-3 rounded-xl bg-amber-400 hover:bg-amber-300 text-slate-950 font-bold text-sm shadow-lg transition"
        >
          Daftar ke {unit.name}
        </a>
      </div>
    </div>
  </section>

  <!-- Content Body -->
  <div class="max-w-6xl mx-auto px-4 sm:px-8 py-16">
    <div class="grid grid-cols-1 lg:grid-cols-12 gap-12">
      <!-- Main Column -->
      <div class="lg:col-span-8 space-y-12">
        <!-- Gambaran Umum -->
        <div class="space-y-4">
          <div class="text-xs font-bold uppercase tracking-wider text-emerald-700 flex items-center gap-1.5">
            <BookOpen class="w-4 h-4" /> Gambaran Umum Program
          </div>
          <h2 class="font-serif-luxury font-bold text-2xl sm:text-3xl text-slate-900">
            Membangun Karakter & Kapasitas Intelektual
          </h2>
          <div class="text-slate-700 leading-relaxed text-base whitespace-pre-line bg-white p-6 sm:p-8 rounded-3xl border border-slate-100 shadow-sm">
            {unit.full_desc}
          </div>
        </div>

        <!-- Kurikulum & Kitab / Pembelajaran -->
        {#if unit.curriculum}
          <div class="space-y-4">
            <div class="text-xs font-bold uppercase tracking-wider text-emerald-700 flex items-center gap-1.5">
              <ClipboardList class="w-4 h-4" /> Kurikulum & Mata Pelajaran / Kitab
            </div>
            <h2 class="font-serif-luxury font-bold text-2xl sm:text-3xl text-slate-900">
              Struktur Pembelajaran
            </h2>
            <div class="bg-emerald-950 text-emerald-50 p-6 sm:p-8 rounded-3xl border border-emerald-900 shadow-lg space-y-4 text-sm sm:text-base leading-relaxed whitespace-pre-line">
              {unit.curriculum}
            </div>
          </div>
        {/if}

        <!-- Persyaratan Masuk -->
        {#if unit.requirements}
          <div class="space-y-4">
            <div class="text-xs font-bold uppercase tracking-wider text-emerald-700 flex items-center gap-1.5">
              <ShieldCheck class="w-4 h-4" /> Persyaratan Pendaftaran
            </div>
            <h2 class="font-serif-luxury font-bold text-2xl sm:text-3xl text-slate-900">
              Syarat & Prosedur Penerimaan
            </h2>
            <div class="bg-white p-6 sm:p-8 rounded-3xl border border-slate-100 shadow-sm space-y-3 text-slate-700 text-sm sm:text-base leading-relaxed whitespace-pre-line">
              {unit.requirements}
            </div>
          </div>
        {/if}
      </div>

      <!-- Sidebar Column -->
      <div class="lg:col-span-4 space-y-8">
        <!-- Pimpinan Unit Card -->
        <div class="bg-white rounded-3xl p-6 border border-slate-100 shadow-md space-y-4 text-center">
          <div class="w-24 h-24 rounded-full overflow-hidden mx-auto bg-slate-100 border-2 border-emerald-600 shadow-md">
            <img
              src={unit.headmaster_photo || 'https://images.unsplash.com/photo-1507003211169-0a1dd7228f2d?auto=format&fit=crop&w=400&q=80'}
              alt={unit.headmaster_name}
              class="w-full h-full object-cover"
            />
          </div>
          <div>
            <div class="text-xs text-slate-400 font-semibold uppercase tracking-wider">
              {unit.type === 'formal' ? 'Kepala Madrasah' : 'Kepala Pengasuh'}
            </div>
            <h3 class="font-serif-luxury font-bold text-lg text-slate-900 mt-0.5">
              {unit.headmaster_name}
            </h3>
          </div>
          <div class="pt-3 border-t border-slate-100">
            <a
              href="/kontak"
              class="w-full py-2.5 px-4 rounded-xl bg-emerald-800 hover:bg-emerald-700 text-white font-semibold text-xs transition block"
            >
              Konsultasi Pendaftaran
            </a>
          </div>
        </div>

        <!-- Fasilitas Khusus Unit -->
        {#if unit.facilities_list}
          <div class="bg-emerald-50/70 rounded-3xl p-6 border border-emerald-100 space-y-3">
            <div class="text-xs font-bold uppercase tracking-wider text-emerald-800 flex items-center gap-1.5">
              <Building2 class="w-4 h-4" /> Sarana Pendukung
            </div>
            <p class="text-xs text-slate-700 leading-relaxed whitespace-pre-line">
              {unit.facilities_list}
            </p>
          </div>
        {/if}

        <!-- Quick Links to other units -->
        <div class="bg-white rounded-3xl p-6 border border-slate-100 shadow-sm space-y-3">
          <div class="text-xs font-bold uppercase tracking-wider text-slate-500">
            Unit Pendidikan Lainnya
          </div>
          <div class="space-y-1.5">
            <a href="/pendidikan/madrasah-diniyah-alhamid" class="block p-2 rounded-xl text-xs font-medium text-slate-700 hover:text-emerald-700 hover:bg-emerald-50 transition">
              Madrasah Diniyah Al-Hamid
            </a>
            <a href="/pendidikan/pmq-alhamid" class="block p-2 rounded-xl text-xs font-medium text-slate-700 hover:text-emerald-700 hover:bg-emerald-50 transition">
              PMQ (Pendidikan Murottilil Qur'an)
            </a>
            <a href="/pendidikan/mts-cintamulya" class="block p-2 rounded-xl text-xs font-medium text-slate-700 hover:text-emerald-700 hover:bg-emerald-50 transition">
              MTs Cintamulya
            </a>
            <a href="/pendidikan/ma-cintamulya" class="block p-2 rounded-xl text-xs font-medium text-slate-700 hover:text-emerald-700 hover:bg-emerald-50 transition">
              MA Cintamulya
            </a>
            <a href="/pendidikan/mai-cintamulya" class="block p-2 rounded-xl text-xs font-medium text-slate-700 hover:text-emerald-700 hover:bg-emerald-50 transition">
              MAI Cintamulya
            </a>
          </div>
        </div>
      </div>
    </div>
  </div>
{/if}
