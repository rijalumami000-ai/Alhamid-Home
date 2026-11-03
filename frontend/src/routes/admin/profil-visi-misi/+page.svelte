<script lang="ts">
  import { onMount } from 'svelte';
  import { api } from '$lib/api';
  import { BookOpen, Target, Save, CheckCircle2, ExternalLink, Sparkles, Image, Compass } from 'lucide-svelte';
  import DriveImageInput from '$lib/components/DriveImageInput.svelte';
  import { updateLocalSettings } from '$lib/stores/settings';

  let activeTab = $state<'profil' | 'visi-misi'>('profil');
  let settings = $state<Record<string, string>>({});
  let loading = $state(true);
  let savedSuccess = $state(false);

  async function loadData() {
    try {
      const list = await api.admin.getSettings();
      const map: Record<string, string> = {};
      for (const item of list) {
        map[item.key] = item.value;
      }
      settings = map;
    } catch (e) {
      console.error(e);
    } finally {
      loading = false;
    }
  }

  async function handleSave(e: SubmitEvent) {
    e.preventDefault();
    savedSuccess = false;
    try {
      await api.admin.updateSettings(settings);
      updateLocalSettings(settings);
      savedSuccess = true;
      setTimeout(() => (savedSuccess = false), 3500);
    } catch (err: any) {
      alert(err.message || 'Gagal menyimpan data');
    }
  }

  onMount(() => {
    loadData();
  });
</script>

<svelte:head>
  <title>Kelola Profil & Visi Misi | CMS Al-Hamid</title>
</svelte:head>

<div class="space-y-6 max-w-4xl">
  <div class="flex flex-col sm:flex-row sm:items-center justify-between gap-4">
    <div>
      <h1 class="font-display font-black text-2xl text-slate-900 tracking-tight">
        Profil, Sejarah, Visi & Misi
      </h1>
      <p class="text-xs text-slate-500 mt-1">
        Kelola narasi resmi kelembagaan, sejarah perjalanan pesantren, visi besar, dan butir-butir misi.
      </p>
    </div>
    <div class="flex items-center gap-2">
      <a
        href="/tentang/profil"
        target="_blank"
        class="inline-flex items-center gap-1.5 px-3.5 py-2 rounded-xl border border-slate-200 text-slate-600 hover:text-slate-900 hover:bg-slate-50 text-xs font-semibold transition"
      >
        <span>Lihat Laman Profil</span>
        <ExternalLink class="w-3.5 h-3.5" />
      </a>
      <a
        href="/tentang/visi-misi"
        target="_blank"
        class="inline-flex items-center gap-1.5 px-3.5 py-2 rounded-xl border border-slate-200 text-slate-600 hover:text-slate-900 hover:bg-slate-50 text-xs font-semibold transition"
      >
        <span>Lihat Laman Visi Misi</span>
        <ExternalLink class="w-3.5 h-3.5" />
      </a>
    </div>
  </div>

  {#if savedSuccess}
    <div class="p-4 rounded-2xl bg-emerald-50 border border-emerald-200 text-emerald-800 text-xs font-bold flex items-center gap-2 animate-in fade-in duration-200">
      <CheckCircle2 class="w-4 h-4 text-emerald-600" />
      <span>Perubahan berhasil disimpan dan langsung disinkronisasi ke website utama!</span>
    </div>
  {/if}

  <!-- Navigation Tabs -->
  <div class="flex items-center gap-2 border-b border-slate-200">
    <button
      onclick={() => (activeTab = 'profil')}
      class="px-5 py-3 text-xs font-bold border-b-2 transition flex items-center gap-2 {activeTab === 'profil'
        ? 'border-teal-700 text-teal-800'
        : 'border-transparent text-slate-500 hover:text-slate-800'}"
    >
      <BookOpen class="w-4 h-4" />
      <span>Profil & Sejarah Pesantren</span>
    </button>
    <button
      onclick={() => (activeTab = 'visi-misi')}
      class="px-5 py-3 text-xs font-bold border-b-2 transition flex items-center gap-2 {activeTab === 'visi-misi'
        ? 'border-teal-700 text-teal-800'
        : 'border-transparent text-slate-500 hover:text-slate-800'}"
    >
      <Target class="w-4 h-4" />
      <span>Visi & Misi Strategis</span>
    </button>
  </div>

  <form onsubmit={handleSave} class="space-y-6">
    <!-- TAB 1: PROFIL & SEJARAH -->
    {#if activeTab === 'profil'}
      <div class="bg-white rounded-3xl p-6 sm:p-8 border border-slate-200 shadow-sm space-y-5">
        <div class="text-xs font-bold uppercase tracking-wider text-teal-800 flex items-center gap-1.5 pb-2 border-b border-slate-100">
          <BookOpen class="w-4 h-4" /> Narasi Profil & Jejak Langkah Pesantren
        </div>

        <div class="space-y-1">
          <label for="p_story" class="text-xs font-bold text-slate-700 uppercase">Narasi Profil Lengkap *</label>
          <textarea
            id="p_story"
            rows="8"
            required
            bind:value={settings['about_profile']}
            placeholder="Tuliskan profil lembaga, latar belakang berdirinya pesantren, komitmen pengajaran turats dan sains..."
            class="w-full px-4 py-3 rounded-xl border border-slate-200 text-sm font-medium leading-relaxed"
          ></textarea>
          <p class="text-[11px] text-slate-400">
            Teks ini akan ditampilkan pada paragraf utama di laman <code>/tentang/profil</code>.
          </p>
        </div>

        <DriveImageInput
          id="p_img"
          label="Foto Suasana / Arsip Sejarah Pesantren"
          bind:value={settings['about_profile_image']}
          placeholder="https://drive.google.com/... atau URL gambar"
          helperText="Foto pendukung yang dipajang di samping narasi sejarah pesantren."
        />
      </div>
    {/if}

    <!-- TAB 2: VISI & MISI -->
    {#if activeTab === 'visi-misi'}
      <div class="bg-white rounded-3xl p-6 sm:p-8 border border-slate-200 shadow-sm space-y-6">
        <div class="text-xs font-bold uppercase tracking-wider text-teal-800 flex items-center gap-1.5 pb-2 border-b border-slate-100">
          <Target class="w-4 h-4" /> Visi & Pernyataan Misi Lembaga
        </div>

        <div class="space-y-1">
          <label for="v_vision" class="text-xs font-bold text-slate-700 uppercase">Visi Utama Pesantren *</label>
          <textarea
            id="v_vision"
            rows="3"
            required
            bind:value={settings['about_vision']}
            placeholder="Terwujudnya lembaga pendidikan Islam unggul yang melahirkan generasi berakhlakul karimah..."
            class="w-full px-4 py-3 rounded-xl border border-slate-200 text-sm font-medium leading-relaxed"
          ></textarea>
        </div>

        <div class="space-y-1">
          <label for="v_mission" class="text-xs font-bold text-slate-700 uppercase">Poin-Poin Misi Lembaga (Satu Baris Per Poin) *</label>
          <textarea
            id="v_mission"
            rows="8"
            required
            bind:value={settings['about_mission']}
            placeholder="Menyelenggarakan kajian keislaman berbasis kutubut turats dengan sanad muttashil.&#10;Mengintegrasikan kurikulum madrasah formal modern (MTs, MA, MAI).&#10;Membina penguasaan bahasa internasional (Arab dan Inggris) aktif.&#10;Menanamkan jiwa santripreneurship dan kepedulian sosial."
            class="w-full px-4 py-3 rounded-xl border border-slate-200 text-sm font-medium leading-relaxed font-mono text-xs"
          ></textarea>
          <p class="text-[11px] text-slate-400">
            Setiap baris baru akan secara otomatis diubah menjadi kotak kartu misi dengan ikon centang di laman <code>/tentang/visi-misi</code>.
          </p>
        </div>
      </div>
    {/if}

    <div class="flex justify-end pt-2">
      <button
        type="submit"
        class="inline-flex items-center gap-2 px-8 py-3 rounded-xl bg-teal-700 hover:bg-teal-800 text-white font-bold text-xs shadow-lg shadow-teal-900/20 transition active:scale-95"
      >
        <Save class="w-4 h-4" />
        <span>Simpan Perubahan</span>
      </button>
    </div>
  </form>
</div>
