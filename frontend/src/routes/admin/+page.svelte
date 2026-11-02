<script lang="ts">
  import { onMount } from 'svelte';
  import { api } from '$lib/api';
  import {
    FileText,
    GraduationCap,
    Image,
    Video,
    Mail,
    Users,
    ArrowUpRight,
    PlusCircle,
    Calendar,
    CheckCircle
  } from 'lucide-svelte';

  let data = $state<any>(null);
  let loading = $state(true);

  onMount(async () => {
    try {
      data = await api.admin.getDashboard();
    } catch (e) {
      console.error(e);
    } finally {
      loading = false;
    }
  });
</script>

<svelte:head>
  <title>Dashboard Ringkasan | CMS Al-Hamid</title>
</svelte:head>

<div class="space-y-8">
  <!-- Page Header -->
  <div class="flex flex-col sm:flex-row sm:items-center justify-between gap-4">
    <div>
      <h1 class="font-serif-luxury font-bold text-2xl sm:text-3xl text-slate-900">
        Dashboard Utama
      </h1>
      <p class="text-xs sm:text-sm text-slate-500 mt-1">
        Ringkasan data operasional website & konten Pondok Pesantren Al-Hamid.
      </p>
    </div>
    <div class="flex items-center gap-2">
      <a
        href="/admin/berita"
        class="inline-flex items-center gap-1.5 px-4 py-2 rounded-xl bg-emerald-800 hover:bg-emerald-700 text-white text-xs font-semibold shadow-sm transition"
      >
        <PlusCircle class="w-4 h-4" /> Tulis Berita Baru
      </a>
      <a
        href="/"
        target="_blank"
        class="inline-flex items-center gap-1.5 px-4 py-2 rounded-xl bg-white border border-slate-200 text-slate-700 text-xs font-semibold hover:bg-slate-50 transition"
      >
        <span>Lihat Website</span>
        <ArrowUpRight class="w-3.5 h-3.5" />
      </a>
    </div>
  </div>

  <!-- Metric Cards -->
  <div class="grid grid-cols-2 md:grid-cols-3 lg:grid-cols-6 gap-4">
    <div class="bg-white p-5 rounded-2xl border border-slate-100 shadow-sm space-y-2">
      <div class="w-8 h-8 rounded-lg bg-emerald-100 text-emerald-800 flex items-center justify-center">
        <FileText class="w-4 h-4" />
      </div>
      <div class="text-2xl font-bold text-slate-900">{data?.stats?.total_articles ?? 0}</div>
      <div class="text-[11px] font-semibold text-slate-500 uppercase">Artikel Berita</div>
    </div>

    <div class="bg-white p-5 rounded-2xl border border-slate-100 shadow-sm space-y-2">
      <div class="w-8 h-8 rounded-lg bg-blue-100 text-blue-800 flex items-center justify-center">
        <GraduationCap class="w-4 h-4" />
      </div>
      <div class="text-2xl font-bold text-slate-900">{data?.stats?.total_units ?? 0}</div>
      <div class="text-[11px] font-semibold text-slate-500 uppercase">Unit Pendidikan</div>
    </div>

    <div class="bg-white p-5 rounded-2xl border border-slate-100 shadow-sm space-y-2">
      <div class="w-8 h-8 rounded-lg bg-purple-100 text-purple-800 flex items-center justify-center">
        <Users class="w-4 h-4" />
      </div>
      <div class="text-2xl font-bold text-slate-900">{data?.stats?.total_caregivers ?? 0}</div>
      <div class="text-[11px] font-semibold text-slate-500 uppercase">Dewan Pengasuh</div>
    </div>

    <div class="bg-white p-5 rounded-2xl border border-slate-100 shadow-sm space-y-2">
      <div class="w-8 h-8 rounded-lg bg-amber-100 text-amber-800 flex items-center justify-center">
        <Image class="w-4 h-4" />
      </div>
      <div class="text-2xl font-bold text-slate-900">{data?.stats?.total_photos ?? 0}</div>
      <div class="text-[11px] font-semibold text-slate-500 uppercase">Foto Galeri</div>
    </div>

    <div class="bg-white p-5 rounded-2xl border border-slate-100 shadow-sm space-y-2">
      <div class="w-8 h-8 rounded-lg bg-red-100 text-red-800 flex items-center justify-center">
        <Video class="w-4 h-4" />
      </div>
      <div class="text-2xl font-bold text-slate-900">{data?.stats?.total_videos ?? 0}</div>
      <div class="text-[11px] font-semibold text-slate-500 uppercase">Video YouTube</div>
    </div>

    <div class="bg-white p-5 rounded-2xl border border-slate-100 shadow-sm space-y-2">
      <div class="w-8 h-8 rounded-lg bg-rose-100 text-rose-800 flex items-center justify-center">
        <Mail class="w-4 h-4" />
      </div>
      <div class="text-2xl font-bold text-slate-900">{data?.stats?.unread_messages ?? 0}</div>
      <div class="text-[11px] font-semibold text-slate-500 uppercase">Pesan Masuk Baru</div>
    </div>
  </div>

  <!-- Content Sections -->
  <div class="grid grid-cols-1 lg:grid-cols-2 gap-8">
    <!-- Recent Articles -->
    <div class="bg-white rounded-3xl p-6 border border-slate-100 shadow-sm space-y-4">
      <div class="flex items-center justify-between pb-3 border-b border-slate-100">
        <div class="font-serif-luxury font-bold text-lg text-slate-900">
          Artikel Terbaru
        </div>
        <a href="/admin/berita" class="text-xs font-semibold text-emerald-700 hover:text-emerald-800">
          Kelola Semua &rarr;
        </a>
      </div>

      <div class="space-y-3">
        {#if data?.recent_articles}
          {#each data.recent_articles as art}
            <div class="flex items-center justify-between p-3 rounded-xl hover:bg-slate-50 transition">
              <div class="space-y-0.5 max-w-[80%]">
                <div class="font-semibold text-xs text-slate-800 line-clamp-1">{art.title}</div>
                <div class="text-[11px] text-slate-400 flex items-center gap-2">
                  <span>{art.category?.name || 'Kajian'}</span>
                  <span>&bull;</span>
                  <span>{art.author}</span>
                </div>
              </div>
              <span class="px-2 py-0.5 rounded-full text-[10px] font-bold uppercase {art.status === 'published' ? 'bg-emerald-100 text-emerald-800' : 'bg-slate-200 text-slate-700'}">
                {art.status}
              </span>
            </div>
          {/each}
        {/if}
      </div>
    </div>

    <!-- Recent Messages -->
    <div class="bg-white rounded-3xl p-6 border border-slate-100 shadow-sm space-y-4">
      <div class="flex items-center justify-between pb-3 border-b border-slate-100">
        <div class="font-serif-luxury font-bold text-lg text-slate-900">
          Pesan Masuk Pengunjung
        </div>
        <a href="/admin/pesan" class="text-xs font-semibold text-emerald-700 hover:text-emerald-800">
          Buka Kotak Masuk &rarr;
        </a>
      </div>

      <div class="space-y-3">
        {#if data?.recent_messages && data.recent_messages.length > 0}
          {#each data.recent_messages as msg}
            <div class="p-3.5 rounded-xl border border-slate-100 space-y-1.5 hover:bg-slate-50 transition">
              <div class="flex items-center justify-between">
                <span class="font-bold text-xs text-slate-900">{msg.name}</span>
                <span class="text-[10px] text-slate-400">
                  {new Date(msg.created_at).toLocaleDateString('id-ID')}
                </span>
              </div>
              <div class="text-xs text-emerald-700 font-medium">{msg.subject}</div>
              <p class="text-xs text-slate-500 line-clamp-2">{msg.message}</p>
            </div>
          {/each}
        {:else}
          <div class="text-center py-8 text-xs text-slate-400">
            Belum ada pesan masuk baru dari formulir kontak.
          </div>
        {/if}
      </div>
    </div>
  </div>
</div>

