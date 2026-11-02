<script lang="ts">
  import { onMount } from 'svelte';
  import { api } from '$lib/api';
  import type { MediaPhoto, MediaVideo, Category } from '$lib/types';
  import { Plus, Trash2, X, Play, Image, Video } from 'lucide-svelte';

  let activeTab = $state<'photos' | 'videos'>('photos');
  let photos = $state<MediaPhoto[]>([]);
  let videos = $state<MediaVideo[]>([]);
  let photoCategories = $state<Category[]>([]);
  let videoCategories = $state<Category[]>([]);

  // Modals
  let isPhotoModalOpen = $state(false);
  let photoForm = $state({
    title: '',
    category_id: 1,
    image_url: '',
    caption: '',
  });

  let isVideoModalOpen = $state(false);
  let videoForm = $state({
    title: '',
    category_id: 1,
    youtube_url: '',
    description: '',
  });

  async function loadData() {
    try {
      const [pList, vList, pCats, vCats] = await Promise.all([
        api.admin.getPhotos(),
        api.admin.getVideos(),
        api.getCategories('galeri_foto'),
        api.getCategories('galeri_video'),
      ]);
      photos = pList;
      videos = vList;
      photoCategories = pCats;
      videoCategories = vCats;
      if (pCats.length > 0) photoForm.category_id = pCats[0].id;
      if (vCats.length > 0) videoForm.category_id = vCats[0].id;
    } catch (e) {
      console.error(e);
    }
  }

  function openCreatePhoto() {
    photoForm = {
      title: '',
      category_id: photoCategories[0]?.id || 1,
      image_url: 'https://images.unsplash.com/photo-1542838132-92c53300491e?auto=format&fit=crop&w=1000&q=80',
      caption: '',
    };
    isPhotoModalOpen = true;
  }

  function openCreateVideo() {
    videoForm = {
      title: '',
      category_id: videoCategories[0]?.id || 1,
      youtube_url: 'https://www.youtube.com/watch?v=dQw4w9WgXcQ',
      description: '',
    };
    isVideoModalOpen = true;
  }

  async function handleSavePhoto(e: SubmitEvent) {
    e.preventDefault();
    try {
      await api.admin.createPhoto(photoForm);
      isPhotoModalOpen = false;
      await loadData();
    } catch (err: any) {
      alert(err.message || 'Gagal menyimpan foto');
    }
  }

  async function handleDeletePhoto(id: number) {
    if (confirm('Hapus foto ini dari galeri?')) {
      try {
        await api.admin.deletePhoto(id);
        await loadData();
      } catch (err: any) {
        alert(err.message || 'Gagal menghapus');
      }
    }
  }

  async function handleSaveVideo(e: SubmitEvent) {
    e.preventDefault();
    try {
      await api.admin.createVideo(videoForm);
      isVideoModalOpen = false;
      await loadData();
    } catch (err: any) {
      alert(err.message || 'Gagal menyimpan video');
    }
  }

  async function handleDeleteVideo(id: number) {
    if (confirm('Hapus video ini dari galeri?')) {
      try {
        await api.admin.deleteVideo(id);
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
  <title>Kelola Galeri Media | CMS Al-Hamid</title>
</svelte:head>

<div class="space-y-6">
  <div class="flex flex-col sm:flex-row sm:items-center justify-between gap-4">
    <div>
      <h1 class="font-serif-luxury font-bold text-2xl text-slate-900">
        Galeri Foto & Video YouTube
      </h1>
      <p class="text-xs text-slate-500 mt-1">
        Kelola album dokumentasi foto kegiatan santri dan sematan video pengajian resmi.
      </p>
    </div>
    <div>
      {#if activeTab === 'photos'}
        <button
          onclick={openCreatePhoto}
          class="inline-flex items-center gap-1.5 px-4 py-2.5 rounded-xl bg-emerald-800 hover:bg-emerald-700 text-white text-xs font-bold shadow-md transition"
        >
          <Plus class="w-4 h-4" /> Tambah Foto
        </button>
      {:else}
        <button
          onclick={openCreateVideo}
          class="inline-flex items-center gap-1.5 px-4 py-2.5 rounded-xl bg-emerald-800 hover:bg-emerald-700 text-white text-xs font-bold shadow-md transition"
        >
          <Plus class="w-4 h-4" /> Tambah Video YouTube
        </button>
      {/if}
    </div>
  </div>

  <!-- Tabs -->
  <div class="flex items-center gap-2 border-b border-slate-200">
    <button
      onclick={() => (activeTab = 'photos')}
      class="px-5 py-3 text-xs font-bold border-b-2 transition {activeTab === 'photos'
        ? 'border-emerald-700 text-emerald-800'
        : 'border-transparent text-slate-500 hover:text-slate-800'}"
    >
      Galeri Foto ({photos.length})
    </button>
    <button
      onclick={() => (activeTab = 'videos')}
      class="px-5 py-3 text-xs font-bold border-b-2 transition {activeTab === 'videos'
        ? 'border-emerald-700 text-emerald-800'
        : 'border-transparent text-slate-500 hover:text-slate-800'}"
    >
      Galeri Video ({videos.length})
    </button>
  </div>

  <!-- PHOTOS TAB -->
  {#if activeTab === 'photos'}
    <div class="grid grid-cols-2 sm:grid-cols-3 md:grid-cols-4 gap-4">
      {#each photos as p}
        <div class="bg-white rounded-2xl border border-slate-200 overflow-hidden shadow-sm flex flex-col justify-between group">
          <div class="relative aspect-square bg-slate-900">
            <img src={p.image_url} alt={p.title} class="w-full h-full object-cover" />
            <span class="absolute top-2 left-2 px-2 py-0.5 rounded bg-slate-950/80 text-white text-[9px] font-bold uppercase">
              {p.category?.name}
            </span>
          </div>
          <div class="p-3 space-y-1">
            <div class="font-bold text-xs text-slate-900 line-clamp-1">{p.title}</div>
            <div class="flex items-center justify-between pt-2 border-t border-slate-100">
              <span class="text-[10px] text-slate-400">
                {new Date(p.created_at).toLocaleDateString('id-ID')}
              </span>
              <button
                onclick={() => handleDeletePhoto(p.id)}
                class="text-red-500 hover:text-red-700 p-1"
                title="Hapus Foto"
              >
                <Trash2 class="w-3.5 h-3.5" />
              </button>
            </div>
          </div>
        </div>
      {/each}
    </div>
  {/if}

  <!-- VIDEOS TAB -->
  {#if activeTab === 'videos'}
    <div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-6">
      {#each videos as v}
        <div class="bg-white rounded-3xl border border-slate-200 overflow-hidden shadow-sm flex flex-col justify-between">
          <div>
            <div class="relative aspect-video bg-slate-900">
              <img src="https://img.youtube.com/vi/{v.youtube_id}/hqdefault.jpg" alt={v.title} class="w-full h-full object-cover" />
              <div class="absolute inset-0 bg-slate-950/30 flex items-center justify-center">
                <Play class="w-10 h-10 text-white fill-white" />
              </div>
              <span class="absolute top-2 left-2 px-2 py-0.5 rounded bg-slate-950/80 text-white text-[9px] font-bold uppercase">
                {v.category?.name}
              </span>
            </div>
            <div class="p-4 space-y-1">
              <h2 class="font-bold text-sm text-slate-900 line-clamp-2">{v.title}</h2>
              <div class="text-xs text-slate-400 font-mono">ID: {v.youtube_id}</div>
            </div>
          </div>

          <div class="p-4 border-t border-slate-100 flex items-center justify-between bg-slate-50/50">
            <a
              href={v.youtube_url}
              target="_blank"
              class="text-xs font-semibold text-emerald-700 hover:underline"
            >
              Buka di YouTube &rarr;
            </a>
            <button
              onclick={() => handleDeleteVideo(v.id)}
              class="p-1.5 rounded-lg text-slate-400 hover:text-red-600 hover:bg-red-50"
            >
              <Trash2 class="w-4 h-4" />
            </button>
          </div>
        </div>
      {/each}
    </div>
  {/if}
</div>

<!-- PHOTO MODAL -->
{#if isPhotoModalOpen}
  <div class="fixed inset-0 z-50 bg-slate-950/60 backdrop-blur-sm flex items-center justify-center p-4">
    <div class="bg-white w-full max-w-md rounded-3xl shadow-2xl border border-slate-200 overflow-hidden">
      <div class="p-6 border-b border-slate-100 flex items-center justify-between">
        <h2 class="font-serif-luxury font-bold text-lg text-slate-900">Tambah Foto ke Galeri</h2>
        <button onclick={() => (isPhotoModalOpen = false)}><X class="w-5 h-5 text-slate-400" /></button>
      </div>
      <form onsubmit={handleSavePhoto} class="p-6 space-y-4">
        <div class="space-y-1">
          <label for="p_title" class="text-xs font-bold text-slate-700 uppercase">Judul Foto *</label>
          <input
            id="p_title"
            type="text"
            required
            bind:value={photoForm.title}
            placeholder="Momen pengajian..."
            class="w-full px-4 py-2.5 rounded-xl border border-slate-200 text-sm font-medium"
          />
        </div>
        <div class="space-y-1">
          <label for="p_cat" class="text-xs font-bold text-slate-700 uppercase">Kategori Album</label>
          <select
            id="p_cat"
            bind:value={photoForm.category_id}
            class="w-full px-4 py-2.5 rounded-xl border border-slate-200 text-sm font-medium bg-white"
          >
            {#each photoCategories as c}
              <option value={c.id}>{c.name}</option>
            {/each}
          </select>
        </div>
        <div class="space-y-1">
          <label for="p_img" class="text-xs font-bold text-slate-700 uppercase">URL Gambar *</label>
          <input
            id="p_img"
            type="url"
            required
            bind:value={photoForm.image_url}
            placeholder="https://images.unsplash.com/..."
            class="w-full px-4 py-2.5 rounded-xl border border-slate-200 text-sm font-medium"
          />
        </div>
        <div class="space-y-1">
          <label for="p_cap" class="text-xs font-bold text-slate-700 uppercase">Keterangan / Kepsyen</label>
          <textarea
            id="p_cap"
            rows="2"
            bind:value={photoForm.caption}
            class="w-full px-4 py-2.5 rounded-xl border border-slate-200 text-sm font-medium"
          ></textarea>
        </div>
        <div class="pt-4 border-t border-slate-100 flex justify-end gap-2">
          <button type="button" onclick={() => (isPhotoModalOpen = false)} class="px-4 py-2 rounded-xl border border-slate-200 text-xs font-bold">Batal</button>
          <button type="submit" class="px-5 py-2 rounded-xl bg-emerald-800 text-white text-xs font-bold shadow-md">Simpan Foto</button>
        </div>
      </form>
    </div>
  </div>
{/if}

<!-- VIDEO MODAL -->
{#if isVideoModalOpen}
  <div class="fixed inset-0 z-50 bg-slate-950/60 backdrop-blur-sm flex items-center justify-center p-4">
    <div class="bg-white w-full max-w-md rounded-3xl shadow-2xl border border-slate-200 overflow-hidden">
      <div class="p-6 border-b border-slate-100 flex items-center justify-between">
        <h2 class="font-serif-luxury font-bold text-lg text-slate-900">Tambah Video YouTube</h2>
        <button onclick={() => (isVideoModalOpen = false)}><X class="w-5 h-5 text-slate-400" /></button>
      </div>
      <form onsubmit={handleSaveVideo} class="p-6 space-y-4">
        <div class="space-y-1">
          <label for="v_title" class="text-xs font-bold text-slate-700 uppercase">Judul Video *</label>
          <input
            id="v_title"
            type="text"
            required
            bind:value={videoForm.title}
            placeholder="Kajian Kitab..."
            class="w-full px-4 py-2.5 rounded-xl border border-slate-200 text-sm font-medium"
          />
        </div>
        <div class="space-y-1">
          <label for="v_cat" class="text-xs font-bold text-slate-700 uppercase">Kategori Video</label>
          <select
            id="v_cat"
            bind:value={videoForm.category_id}
            class="w-full px-4 py-2.5 rounded-xl border border-slate-200 text-sm font-medium bg-white"
          >
            {#each videoCategories as c}
              <option value={c.id}>{c.name}</option>
            {/each}
          </select>
        </div>
        <div class="space-y-1">
          <label for="v_url" class="text-xs font-bold text-slate-700 uppercase">Tautan / URL YouTube *</label>
          <input
            id="v_url"
            type="url"
            required
            bind:value={videoForm.youtube_url}
            placeholder="https://www.youtube.com/watch?v=..."
            class="w-full px-4 py-2.5 rounded-xl border border-slate-200 text-sm font-medium"
          />
        </div>
        <div class="space-y-1">
          <label for="v_desc" class="text-xs font-bold text-slate-700 uppercase">Deskripsi Singkat</label>
          <textarea
            id="v_desc"
            rows="2"
            bind:value={videoForm.description}
            class="w-full px-4 py-2.5 rounded-xl border border-slate-200 text-sm font-medium"
          ></textarea>
        </div>
        <div class="pt-4 border-t border-slate-100 flex justify-end gap-2">
          <button type="button" onclick={() => (isVideoModalOpen = false)} class="px-4 py-2 rounded-xl border border-slate-200 text-xs font-bold">Batal</button>
          <button type="submit" class="px-5 py-2 rounded-xl bg-emerald-800 text-white text-xs font-bold shadow-md">Simpan Video</button>
        </div>
      </form>
    </div>
  </div>
{/if}

