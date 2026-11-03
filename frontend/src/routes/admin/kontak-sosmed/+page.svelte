<script lang="ts">
  import { onMount } from 'svelte';
  import { api } from '$lib/api';
  import { MapPin, Phone, Mail, Share2, Save, CheckCircle2, ExternalLink, Globe } from 'lucide-svelte';
  import { updateLocalSettings } from '$lib/stores/settings';

  let settings = $state<Record<string, string>>({
    site_address: '',
    site_phone: '',
    site_whatsapp: '',
    site_email: '',
    site_office_hours: '',
    site_maps: '',
    site_maps_embed: '',
    social_facebook: '',
    social_instagram: '',
    social_youtube: '',
    social_tiktok: ''
  });
  let loading = $state(true);
  let savedSuccess = $state(false);

  async function loadData() {
    try {
      const list = await api.admin.getSettings();
      const map: Record<string, string> = {
        site_address: '',
        site_phone: '',
        site_whatsapp: '',
        site_email: '',
        site_office_hours: '',
        site_maps: '',
        site_maps_embed: '',
        social_facebook: '',
        social_instagram: '',
        social_youtube: '',
        social_tiktok: ''
      };
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
      alert(err.message || 'Gagal menyimpan pengaturan kontak');
    }
  }

  onMount(() => {
    loadData();
  });
</script>

<svelte:head>
  <title>Kontak, Lokasi & Media Sosial | CMS Al-Hamid</title>
</svelte:head>

<div class="space-y-6 max-w-4xl">
  <div class="flex flex-col sm:flex-row sm:items-center justify-between gap-4">
    <div>
      <h1 class="font-display font-black text-2xl text-slate-900 tracking-tight">
        Kontak, Lokasi & Media Sosial
      </h1>
      <p class="text-xs text-slate-500 mt-1">
        Atur alamat sekretariat, tautan Google Maps, nomor telepon, WhatsApp, dan tautan jejaring sosial resmi.
      </p>
    </div>
    <a
      href="/kontak"
      target="_blank"
      class="inline-flex items-center gap-1.5 px-3.5 py-2 rounded-xl border border-slate-200 text-slate-600 hover:text-slate-900 hover:bg-slate-50 text-xs font-semibold transition"
    >
      <span>Lihat Laman Kontak</span>
      <ExternalLink class="w-3.5 h-3.5" />
    </a>
  </div>

  {#if savedSuccess}
    <div class="p-4 rounded-2xl bg-emerald-50 border border-emerald-200 text-emerald-800 text-xs font-bold flex items-center gap-2 animate-in fade-in duration-200">
      <CheckCircle2 class="w-4 h-4 text-emerald-600" />
      <span>Informasi kontak dan media sosial berhasil diperbarui dan langsung tayang di web utama!</span>
    </div>
  {/if}

  {#if loading}
    <div class="p-12 text-center text-slate-400 bg-white rounded-3xl border border-slate-200 shadow-sm">
      <div class="animate-spin w-8 h-8 border-4 border-teal-500 border-t-transparent rounded-full mx-auto mb-3"></div>
      <p class="text-xs font-semibold">Memuat data kontak & media sosial...</p>
    </div>
  {:else}
    <form onsubmit={handleSave} class="space-y-6">
    <!-- SECTION 1: LOKASI & GOOGLE MAPS -->
    <div class="bg-white rounded-3xl p-6 sm:p-8 border border-slate-200 shadow-sm space-y-4">
      <div class="text-xs font-bold uppercase tracking-wider text-teal-800 flex items-center gap-1.5 pb-2 border-b border-slate-100">
        <MapPin class="w-4 h-4" /> Alamat Fisik & Tautan Google Maps
      </div>

      <div class="space-y-1">
        <label for="c_addr" class="text-xs font-bold text-slate-700 uppercase">Alamat Lengkap Kampus / Sekretariat *</label>
        <textarea
          id="c_addr"
          rows="2"
          required
          bind:value={settings['site_address']}
          placeholder="Jl. Pesantren Al-Hamid No. 01, Cintamulya, Candipuro, Lampung Selatan, Lampung 35356"
          class="w-full px-4 py-2.5 rounded-xl border border-slate-200 text-sm font-medium leading-relaxed"
        ></textarea>
      </div>

      <div class="space-y-1">
        <label for="c_maps" class="text-xs font-bold text-slate-700 uppercase">Tautan Link Google Maps *</label>
        <input
          id="c_maps"
          type="url"
          required
          bind:value={settings['site_maps']}
          placeholder="https://maps.google.com/?q=Pondok+Pesantren+Al-Hamid"
          class="w-full px-4 py-2.5 rounded-xl border border-slate-200 text-sm font-mono text-slate-800"
        />
        <p class="text-[11px] text-slate-400">
          Tautan ini akan dibuka saat pengunjung mengklik tombol <em>"Buka di Google Maps"</em> pada laman kontak.
        </p>
      </div>
    </div>

    <!-- SECTION 2: TELEPON, WHATSAPP & EMAIL -->
    <div class="bg-white rounded-3xl p-6 sm:p-8 border border-slate-200 shadow-sm space-y-4">
      <div class="text-xs font-bold uppercase tracking-wider text-teal-800 flex items-center gap-1.5 pb-2 border-b border-slate-100">
        <Phone class="w-4 h-4" /> Saluran Komunikasi & Pelayanan
      </div>

      <div class="grid grid-cols-1 sm:grid-cols-2 gap-4">
        <div class="space-y-1">
          <label for="c_wa" class="text-xs font-bold text-slate-700 uppercase">Nomor WhatsApp Resmi (Kode 62...)</label>
          <input
            id="c_wa"
            type="text"
            bind:value={settings['site_whatsapp']}
            placeholder="6281234567890"
            class="w-full px-4 py-2.5 rounded-xl border border-slate-200 text-sm font-mono"
          />
          <p class="text-[11px] text-slate-400">Gunakan angka saja tanpa spasi/tanda hubung, diawali 62.</p>
        </div>

        <div class="space-y-1">
          <label for="c_phone" class="text-xs font-bold text-slate-700 uppercase">Nomor Telepon Kantor / Suara</label>
          <input
            id="c_phone"
            type="text"
            bind:value={settings['site_phone']}
            placeholder="+62 812-3456-7890"
            class="w-full px-4 py-2.5 rounded-xl border border-slate-200 text-sm font-medium"
          />
        </div>
      </div>

      <div class="grid grid-cols-1 sm:grid-cols-2 gap-4">
        <div class="space-y-1">
          <label for="c_email" class="text-xs font-bold text-slate-700 uppercase">Email Resmi Sekretariat</label>
          <input
            id="c_email"
            type="email"
            bind:value={settings['site_email']}
            placeholder="sekretariat@alhamid.sch.id"
            class="w-full px-4 py-2.5 rounded-xl border border-slate-200 text-sm font-medium"
          />
        </div>

        <div class="space-y-1">
          <label for="c_hours" class="text-xs font-bold text-slate-700 uppercase">Jam Pelayanan Kantor</label>
          <input
            id="c_hours"
            type="text"
            bind:value={settings['site_office_hours']}
            placeholder="Setiap Hari: 08.00 – 16.30 WIB"
            class="w-full px-4 py-2.5 rounded-xl border border-slate-200 text-sm font-medium"
          />
        </div>
      </div>
    </div>

    <!-- SECTION 3: MEDIA SOSIAL -->
    <div class="bg-white rounded-3xl p-6 sm:p-8 border border-slate-200 shadow-sm space-y-4">
      <div class="text-xs font-bold uppercase tracking-wider text-teal-800 flex items-center gap-1.5 pb-2 border-b border-slate-100">
        <Share2 class="w-4 h-4" /> Akun & Jejaring Media Sosial Resmi
      </div>

      <div class="grid grid-cols-1 sm:grid-cols-2 gap-4">
        <div class="space-y-1">
          <label for="soc_ig" class="text-xs font-bold text-slate-700 uppercase">Instagram URL</label>
          <input
            id="soc_ig"
            type="url"
            bind:value={settings['social_instagram']}
            placeholder="https://instagram.com/alhamid"
            class="w-full px-4 py-2.5 rounded-xl border border-slate-200 text-sm font-mono text-slate-800"
          />
        </div>

        <div class="space-y-1">
          <label for="soc_yt" class="text-xs font-bold text-slate-700 uppercase">YouTube Channel URL</label>
          <input
            id="soc_yt"
            type="url"
            bind:value={settings['social_youtube']}
            placeholder="https://youtube.com/@alhamid"
            class="w-full px-4 py-2.5 rounded-xl border border-slate-200 text-sm font-mono text-slate-800"
          />
        </div>
      </div>

      <div class="grid grid-cols-1 sm:grid-cols-2 gap-4">
        <div class="space-y-1">
          <label for="soc_fb" class="text-xs font-bold text-slate-700 uppercase">Facebook Page URL</label>
          <input
            id="soc_fb"
            type="url"
            bind:value={settings['social_facebook']}
            placeholder="https://facebook.com/alhamid"
            class="w-full px-4 py-2.5 rounded-xl border border-slate-200 text-sm font-mono text-slate-800"
          />
        </div>

        <div class="space-y-1">
          <label for="soc_tt" class="text-xs font-bold text-slate-700 uppercase">TikTok URL</label>
          <input
            id="soc_tt"
            type="url"
            bind:value={settings['social_tiktok']}
            placeholder="https://tiktok.com/@alhamid"
            class="w-full px-4 py-2.5 rounded-xl border border-slate-200 text-sm font-mono text-slate-800"
          />
        </div>
      </div>
    </div>

    <div class="flex justify-end pt-2">
      <button
        type="submit"
        class="inline-flex items-center gap-2 px-8 py-3 rounded-xl bg-teal-700 hover:bg-teal-800 text-white font-bold text-xs shadow-lg shadow-teal-900/20 transition active:scale-95"
      >
        <Save class="w-4 h-4" />
        <span>Simpan Kontak & Sosmed</span>
      </button>
    </div>
  </form>
  {/if}
</div>

