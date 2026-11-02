<script lang="ts">
  import { onMount } from 'svelte';
  import { api } from '$lib/api';
  import type { Banner } from '$lib/types';
  import { Plus, Edit2, Trash2, X, Sliders, ExternalLink } from 'lucide-svelte';

  let list = $state<Banner[]>([]);
  let isModalOpen = $state(false);
  let editing = $state<Banner | null>(null);
  let form = $state({
    title: '',
    subtitle: '',
    image_url: '',
    button_text: 'Jelajahi',
    button_link: '#tentang',
    badge: 'Tahun Ajaran 2026/2027',
    order: 0,
    is_active: true,
  });

  async function loadData() {
    try {
      list = await api.admin.getBanners();
    } catch (e) {
      console.error(e);
    }
  }

  function openCreate() {
    editing = null;
    form = {
      title: '',
      subtitle: '',
      image_url: 'https://images.unsplash.com/photo-1542838132-92c53300491e?auto=format&fit=crop&w=1920&q=80',
      button_text: 'Daftar Sekarang',
      button_link: '/kontak',
      badge: 'Penerimaan Santri Baru',
      order: list.length + 1,
      is_active: true,
    };
    isModalOpen = true;
  }

  function openEdit(b: Banner) {
    editing = b;
    form = {
      title: b.title,
      subtitle: b.subtitle,
      image_url: b.image_url,
      button_text: b.button_text,
      button_link: b.button_link,
      badge: b.badge,
      order: b.order,
      is_active: b.is_active,
    };
    isModalOpen = true;
  }

  async function handleSave(e: SubmitEvent) {
    e.preventDefault();
    try {
      if (editing) {
        await api.admin.updateBanner(editing.id, form);
      } else {
        await api.admin.createBanner(form);
      }
      isModalOpen = false;
      await loadData();
    } catch (err: any) {
      alert(err.message || 'Gagal menyimpan banner');
    }
  }

  async function handleDelete(id: number) {
    if (confirm('Hapus slide banner ini?')) {
      try {
        await api.admin.deleteBanner(id);
        await loadData();
      } catch (err: any) {
        alert(err.message || 'Gagal menghapus');
      }
    }
  }

  onMount(() => {
    loadData();
  });
</script>

<svelte:head>
  <title>Kelola Banner Beranda | CMS Al-Hamid</title>
</svelte:head>

<div class="space-y-6">
  <div class="flex flex-col sm:flex-row sm:items-center justify-between gap-4">
    <div>
      <h1 class="font-serif-luxury font-bold text-2xl text-slate-900">
        Hero Banner & Slider Beranda
      </h1>
      <p class="text-xs text-slate-500 mt-1">
        Atur slide promosi, latar belakang foto, judul besar, dan tombol ajakan aksi di beranda utama.
      </p>
    </div>
    <button
      onclick={openCreate}
      class="inline-flex items-center gap-1.5 px-4 py-2.5 rounded-xl bg-emerald-800 hover:bg-emerald-700 text-white text-xs font-bold shadow-md transition"
    >
      <Plus class="w-4 h-4" /> Tambah Banner Slide
    </button>
  </div>

  <div class="space-y-4">
    {#each list as item}
      <div class="bg-white rounded-3xl border border-slate-200 overflow-hidden shadow-sm flex flex-col md:flex-row gap-6 p-6 items-center">
        <div class="w-full md:w-64 h-36 rounded-2xl overflow-hidden bg-slate-900 shrink-0 relative">
          <img src={item.image_url} alt={item.title} class="w-full h-full object-cover" />
          {#if item.badge}
            <span class="absolute top-2 left-2 px-2 py-0.5 rounded bg-emerald-800 text-white text-[9px] font-bold">
              {item.badge}
            </span>
          {/if}
        </div>

        <div class="flex-1 space-y-2">
          <div class="flex items-center gap-2">
            <span class="px-2 py-0.5 rounded text-[10px] font-bold uppercase {item.is_active ? 'bg-emerald-100 text-emerald-800' : 'bg-slate-200 text-slate-600'}">
              {item.is_active ? 'Aktif Tayang' : 'Nonaktif'}
            </span>
            <span class="text-xs text-slate-400 font-mono">Urutan #{item.order}</span>
          </div>

          <h2 class="font-serif-luxury font-bold text-lg text-slate-900">{item.title}</h2>
          <p class="text-xs text-slate-500 line-clamp-2">{item.subtitle}</p>

          <div class="text-xs text-emerald-700 font-semibold pt-1">
            Tombol: "{item.button_text}" &rarr; <span class="font-mono text-slate-500">{item.button_link}</span>
          </div>
        </div>

        <div class="flex md:flex-col items-center gap-2 shrink-0">
          <button
            onclick={() => openEdit(item)}
            class="p-2 rounded-xl text-slate-600 hover:text-emerald-700 hover:bg-emerald-50 transition"
          >
            <Edit2 class="w-4 h-4" />
          </button>
          <button
            onclick={() => handleDelete(item.id)}
            class="p-2 rounded-xl text-slate-600 hover:text-red-600 hover:bg-red-50 transition"
          >
            <Trash2 class="w-4 h-4" />
          </button>
        </div>
      </div>
    {/each}
  </div>
</div>

<!-- MODAL FORM -->
{#if isModalOpen}
  <div class="fixed inset-0 z-50 bg-slate-950/60 backdrop-blur-sm flex items-center justify-center p-4">
    <div class="bg-white w-full max-w-lg rounded-3xl shadow-2xl border border-slate-200 overflow-hidden">
      <div class="p-6 border-b border-slate-100 flex items-center justify-between">
        <h2 class="font-serif-luxury font-bold text-lg text-slate-900">
          {editing ? 'Edit Banner Slide' : 'Tambah Banner Slide'}
        </h2>
        <button onclick={() => (isModalOpen = false)}><X class="w-5 h-5 text-slate-400" /></button>
      </div>

      <form onsubmit={handleSave} class="p-6 space-y-4 max-h-[80vh] overflow-y-auto">
        <div class="space-y-1">
          <label for="b_title" class="text-xs font-bold text-slate-700 uppercase">Judul Utama Slide *</label>
          <input
            id="b_title"
            type="text"
            required
            bind:value={form.title}
            class="w-full px-4 py-2.5 rounded-xl border border-slate-200 text-sm font-medium"
          />
        </div>

        <div class="space-y-1">
          <label for="b_sub" class="text-xs font-bold text-slate-700 uppercase">Subjudul / Deskripsi Pendukung</label>
          <textarea
            id="b_sub"
            rows="2"
            bind:value={form.subtitle}
            class="w-full px-4 py-2.5 rounded-xl border border-slate-200 text-sm font-medium"
          ></textarea>
        </div>

        <div class="space-y-1">
          <label for="b_img" class="text-xs font-bold text-slate-700 uppercase">URL Background Foto *</label>
          <input
            id="b_img"
            type="url"
            required
            bind:value={form.image_url}
            class="w-full px-4 py-2.5 rounded-xl border border-slate-200 text-sm font-medium"
          />
        </div>

        <div class="grid grid-cols-2 gap-3">
          <div class="space-y-1">
            <label for="b_badge" class="text-xs font-bold text-slate-700 uppercase">Badge Label</label>
            <input
              id="b_badge"
              type="text"
              bind:value={form.badge}
              placeholder="PSB 2026/2027"
              class="w-full px-4 py-2.5 rounded-xl border border-slate-200 text-sm font-medium"
            />
          </div>
          <div class="space-y-1">
            <label for="b_order" class="text-xs font-bold text-slate-700 uppercase">Urutan Tampil</label>
            <input
              id="b_order"
              type="number"
              bind:value={form.order}
              class="w-full px-4 py-2.5 rounded-xl border border-slate-200 text-sm font-medium"
            />
          </div>
        </div>

        <div class="grid grid-cols-2 gap-3">
          <div class="space-y-1">
            <label for="b_btn" class="text-xs font-bold text-slate-700 uppercase">Teks Tombol</label>
            <input
              id="b_btn"
              type="text"
              bind:value={form.button_text}
              placeholder="Daftar Sekarang"
              class="w-full px-4 py-2.5 rounded-xl border border-slate-200 text-sm font-medium"
            />
          </div>
          <div class="space-y-1">
            <label for="b_link" class="text-xs font-bold text-slate-700 uppercase">Link Tombol</label>
            <input
              id="b_link"
              type="text"
              bind:value={form.button_link}
              placeholder="/kontak"
              class="w-full px-4 py-2.5 rounded-xl border border-slate-200 text-sm font-medium"
            />
          </div>
        </div>

        <div class="flex items-center gap-2 pt-2">
          <input id="b_active" type="checkbox" bind:checked={form.is_active} class="w-4 h-4 text-emerald-600 rounded" />
          <label for="b_active" class="text-xs font-bold text-slate-700">Aktifkan Slide Ini di Beranda</label>
        </div>

        <div class="pt-4 border-t border-slate-100 flex justify-end gap-2">
          <button type="button" onclick={() => (isModalOpen = false)} class="px-4 py-2 rounded-xl border border-slate-200 text-xs font-bold">Batal</button>
          <button type="submit" class="px-5 py-2 rounded-xl bg-emerald-800 text-white text-xs font-bold shadow-md">Simpan Slide</button>
        </div>
      </form>
    </div>
  </div>
{/if}

