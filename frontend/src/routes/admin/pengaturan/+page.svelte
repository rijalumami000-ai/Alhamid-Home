<script lang="ts">
  import { onMount } from 'svelte';
  import { api } from '$lib/api';
  import { Settings, Save, CheckCircle2, Globe, Phone, Share2, BookOpen } from 'lucide-svelte';

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
      savedSuccess = true;
      setTimeout(() => (savedSuccess = false), 3500);
    } catch (err: any) {
      alert(err.message || 'Gagal menyimpan pengaturan');
    }
  }

  onMount(() => {
    loadData();
  });
</script>

<svelte:head>
  <title>Pengaturan Website | CMS Al-Hamid</title>
</svelte:head>

<div class="space-y-6 max-w-4xl">
  <div>
    <h1 class="font-serif-luxury font-bold text-2xl text-slate-900">
      Pengaturan Website Pesantren
    </h1>
    <p class="text-xs text-slate-500 mt-1">
      Atur identitas lembaga, kontak resmi, tautan media sosial, serta narasi profil dan visi-misi.
    </p>
  </div>

  {#if savedSuccess}
    <div class="p-4 rounded-2xl bg-emerald-50 border border-emerald-200 text-emerald-800 text-xs font-bold flex items-center gap-2">
      <CheckCircle2 class="w-4 h-4 text-emerald-600" />
      <span>Pengaturan website berhasil diperbarui dan langsung diterapkan ke halaman publik!</span>
    </div>
  {/if}

  <form onsubmit={handleSave} class="space-y-8">
    <!-- SECTION 1: IDENTITAS -->
    <div class="bg-white rounded-3xl p-6 sm:p-8 border border-slate-200 shadow-sm space-y-4">
      <div class="text-xs font-bold uppercase tracking-wider text-emerald-800 flex items-center gap-1.5 pb-2 border-b border-slate-100">
        <Globe class="w-4 h-4" /> Identitas Lembaga & Portal
      </div>

      <div class="space-y-1">
        <label for="s_name" class="text-xs font-bold text-slate-700 uppercase">Nama Pesantren *</label>
        <input
          id="s_name"
          type="text"
          bind:value={settings['site_name']}
          class="w-full px-4 py-2.5 rounded-xl border border-slate-200 text-sm font-medium"
        />
      </div>

      <div class="space-y-1">
        <label for="s_tagline" class="text-xs font-bold text-slate-700 uppercase">Tagline / Slogan</label>
        <input
          id="s_tagline"
          type="text"
          bind:value={settings['site_tagline']}
          class="w-full px-4 py-2.5 rounded-xl border border-slate-200 text-sm font-medium"
        />
      </div>

      <div class="space-y-1">
        <label for="s_desc" class="text-xs font-bold text-slate-700 uppercase">Deskripsi Ringkas (Meta SEO)</label>
        <textarea
          id="s_desc"
          rows="2"
          bind:value={settings['site_description']}
          class="w-full px-4 py-2.5 rounded-xl border border-slate-200 text-sm font-medium"
        ></textarea>
      </div>
    </div>

    <!-- SECTION 2: KONTAK & ALAMAT -->
    <div class="bg-white rounded-3xl p-6 sm:p-8 border border-slate-200 shadow-sm space-y-4">
      <div class="text-xs font-bold uppercase tracking-wider text-emerald-800 flex items-center gap-1.5 pb-2 border-b border-slate-100">
        <Phone class="w-4 h-4" /> Kontak & Lokasi Kampus
      </div>

      <div class="grid grid-cols-1 sm:grid-cols-2 gap-4">
        <div class="space-y-1">
          <label for="s_phone" class="text-xs font-bold text-slate-700 uppercase">Nomor Telepon</label>
          <input
            id="s_phone"
            type="text"
            bind:value={settings['site_phone']}
            class="w-full px-4 py-2.5 rounded-xl border border-slate-200 text-sm font-medium"
          />
        </div>
        <div class="space-y-1">
          <label for="s_wa" class="text-xs font-bold text-slate-700 uppercase">Nomor WhatsApp (Angka saja: 628...)</label>
          <input
            id="s_wa"
            type="text"
            bind:value={settings['site_whatsapp']}
            class="w-full px-4 py-2.5 rounded-xl border border-slate-200 text-sm font-medium"
          />
        </div>
      </div>

      <div class="space-y-1">
        <label for="s_email" class="text-xs font-bold text-slate-700 uppercase">Email Resmi</label>
        <input
          id="s_email"
          type="email"
          bind:value={settings['site_email']}
          class="w-full px-4 py-2.5 rounded-xl border border-slate-200 text-sm font-medium"
        />
      </div>

      <div class="space-y-1">
        <label for="s_addr" class="text-xs font-bold text-slate-700 uppercase">Alamat Lengkap</label>
        <textarea
          id="s_addr"
          rows="2"
          bind:value={settings['site_address']}
          class="w-full px-4 py-2.5 rounded-xl border border-slate-200 text-sm font-medium"
        ></textarea>
      </div>
    </div>

    <!-- SECTION 3: MEDIA SOSIAL -->
    <div class="bg-white rounded-3xl p-6 sm:p-8 border border-slate-200 shadow-sm space-y-4">
      <div class="text-xs font-bold uppercase tracking-wider text-emerald-800 flex items-center gap-1.5 pb-2 border-b border-slate-100">
        <Share2 class="w-4 h-4" /> Tautan Media Sosial
      </div>

      <div class="grid grid-cols-1 sm:grid-cols-2 gap-4">
        <div class="space-y-1">
          <label for="s_ig" class="text-xs font-bold text-slate-700 uppercase">URL Instagram</label>
          <input id="s_ig" type="url" bind:value={settings['social_instagram']} class="w-full px-4 py-2.5 rounded-xl border border-slate-200 text-sm font-medium" />
        </div>
        <div class="space-y-1">
          <label for="s_yt" class="text-xs font-bold text-slate-700 uppercase">URL YouTube</label>
          <input id="s_yt" type="url" bind:value={settings['social_youtube']} class="w-full px-4 py-2.5 rounded-xl border border-slate-200 text-sm font-medium" />
        </div>
        <div class="space-y-1">
          <label for="s_fb" class="text-xs font-bold text-slate-700 uppercase">URL Facebook</label>
          <input id="s_fb" type="url" bind:value={settings['social_facebook']} class="w-full px-4 py-2.5 rounded-xl border border-slate-200 text-sm font-medium" />
        </div>
        <div class="space-y-1">
          <label for="s_tt" class="text-xs font-bold text-slate-700 uppercase">URL TikTok</label>
          <input id="s_tt" type="url" bind:value={settings['social_tiktok']} class="w-full px-4 py-2.5 rounded-xl border border-slate-200 text-sm font-medium" />
        </div>
      </div>
    </div>

    <!-- SECTION 4: PROFIL, VISI, MISI -->
    <div class="bg-white rounded-3xl p-6 sm:p-8 border border-slate-200 shadow-sm space-y-4">
      <div class="text-xs font-bold uppercase tracking-wider text-emerald-800 flex items-center gap-1.5 pb-2 border-b border-slate-100">
        <BookOpen class="w-4 h-4" /> Narasi Profil, Visi & Misi
      </div>

      <div class="space-y-1">
        <label for="s_profile" class="text-xs font-bold text-slate-700 uppercase">Profil & Sejarah Pesantren</label>
        <textarea
          id="s_profile"
          rows="5"
          bind:value={settings['about_profile']}
          class="w-full px-4 py-2.5 rounded-xl border border-slate-200 text-sm font-medium leading-relaxed"
        ></textarea>
      </div>

      <div class="space-y-1">
        <label for="s_vision" class="text-xs font-bold text-slate-700 uppercase">Visi Pesantren</label>
        <textarea
          id="s_vision"
          rows="2"
          bind:value={settings['about_vision']}
          class="w-full px-4 py-2.5 rounded-xl border border-slate-200 text-sm font-medium leading-relaxed"
        ></textarea>
      </div>

      <div class="space-y-1">
        <label for="s_mission" class="text-xs font-bold text-slate-700 uppercase">Misi Pesantren (Daftar Poin)</label>
        <textarea
          id="s_mission"
          rows="5"
          bind:value={settings['about_mission']}
          class="w-full px-4 py-2.5 rounded-xl border border-slate-200 text-sm font-medium leading-relaxed"
        ></textarea>
      </div>
    </div>

    <!-- Submit Button -->
    <div class="flex justify-end">
      <button
        type="submit"
        class="px-8 py-3 rounded-xl bg-emerald-800 hover:bg-emerald-700 text-white font-bold text-sm shadow-lg transition flex items-center gap-2"
      >
        <Save class="w-4 h-4" />
        <span>Simpan Seluruh Pengaturan</span>
      </button>
    </div>
  </form>
</div>

